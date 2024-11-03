package repository

import (
	"errors"
	"skripsi/features/user/entity"
	"skripsi/features/user/interfaces"
	"skripsi/features/user/mapping"
	"skripsi/features/user/model"
	"skripsi/utils/constant"
	"skripsi/utils/pagination"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// FindByEmail implements interfaces.UserRepositoryInterface.
func (ur *userRepository) FindByEmail(email string) (entity.UsersCore, error) {
	dataUser := model.Users{}

	tx := ur.db.Where("email = ?", email).First(&dataUser)

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUser)
	return dataResponse, nil
}

// GetAll implements interfaces.UserRepositoryInterface.
func (ur *userRepository) GetAll(search string, page, limit int) ([]entity.UsersCore, pagination.PageInfo, int, error) {
	dataUser := []model.Users{}

	offset := (page - 1) * limit
	query := ur.db.Model(&model.Users{})

	if search != "" {
		query = query.Where("name LIKE ? or email LIKE ? or username LIKE?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&dataUser)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&dataUser)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	dataResponse := mapping.ListUserModelToUserCore(dataUser)
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataResponse, pageInfo, int(totalCount), nil
}

// GetById implements interfaces.UserRepositoryInterface.
func (ur *userRepository) GetById(id string) (entity.UsersCore, error) {
	panic("unimplemented")
}

// Register implements interfaces.UserRepositoryInterface.
func (ur *userRepository) Register(data entity.UsersCore) (entity.UsersCore, error) {
	request := mapping.UserCoreToUserModel(data)

	tx := ur.db.Create(&request)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(request)
	return dataResponse, nil
}

// UpdateById implements interfaces.UserRepositoryInterface.
func (ur *userRepository) UpdateById(id string, data entity.UsersCore) error {
	panic("unimplemented")
}

// FindByUsername implements interfaces.UserRepositoryInterface.
func (ur *userRepository) FindByUsername(username string) (entity.UsersCore, error) {
	dataUser := model.Users{}

	tx := ur.db.Where("username = ?", username).First(&dataUser)

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constant.ERROR_USERNAME_EXIST)
	}

	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUser)
	return dataResponse, nil
}
