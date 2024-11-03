package repository

import (
	"errors"
	"skripsi/features/user/entity"
	"skripsi/features/user/interfaces"
	"skripsi/features/user/mapping"
	"skripsi/features/user/model"
	"skripsi/utils/constant"

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
func (ur *userRepository) GetAll() ([]entity.UsersCore, error) {
	panic("unimplemented")
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
