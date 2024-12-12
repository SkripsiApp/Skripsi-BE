package repository

import (
	"errors"
	"skripsi/features/user/entity"
	"skripsi/features/user/interfaces"
	"skripsi/features/user/mapping"
	"skripsi/features/user/model"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"
	"time"

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
	dataUser := model.Users{}

	tx := ur.db.Where("id = ?", id).First(&dataUser)

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constant.ERROR_DATA_ID)
	}

	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUser)
	return dataResponse, nil
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
	request := mapping.UserCoreToUserModel(data)

	tx := ur.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constant.ERROR_DATA_ID)
	}

	return nil
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

// SendOTP implements interfaces.UserRepositoryInterface.
func (ur *userRepository) SendOTP(email, otp string, expiry int64) (entity.UsersCore, error) {
	dataUser := model.Users{}

	tx := ur.db.Where("email = ?", email).First(&dataUser)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, helper.ResponseError(404, constant.ERROR_DATA_EMAIL)
		}
		return entity.UsersCore{}, tx.Error
	}

	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataUser.Otp = otp
	dataUser.OtpExpired = expiry

	tx = ur.db.Save(&dataUser)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUser)
	return dataResponse, nil
}

// VerifyOTP implements interfaces.UserRepositoryInterface.
func (ur *userRepository) VerifyOTP(email, otp string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("otp = ? AND email = ?", otp, email).First(&dataUsers)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, helper.ResponseError(404, "email atau otp tidak ditemukan")
		}
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUsers)
	return dataResponse, nil
}

// ResetOTP implements interfaces.UserRepositoryInterface.
func (ur *userRepository) ResetOTP(otp string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("otp = ?", otp).First(&dataUsers)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, helper.ResponseError(404, "otp tidak ditemukan")
		}
		return entity.UsersCore{}, tx.Error
	}

	dataUsers.Otp = ""
	dataUsers.OtpExpired = 0

	tx = ur.db.Save(&dataUsers)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUsers)
	return dataResponse, nil
}

// NewPassword implements interfaces.UserRepositoryInterface.
func (ur *userRepository) NewPassword(email string, data entity.UsersCore) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("email = ?", email).First(&dataUsers)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, helper.ResponseError(404, constant.ERROR_DATA_EMAIL)
		}
		return entity.UsersCore{}, tx.Error
	}

	errUpdate := ur.db.Model(&dataUsers).Updates(mapping.UserCoreToUserModel(data))
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate.Error
	}

	dataResponse := mapping.UserModelToUserCore(dataUsers)

	return dataResponse, nil
}

// updatedPoint implements interfaces.UserRepositoryInterface.
func (ur *userRepository) UpdatedPoint(id string, point int) error {
	tx := ur.db.Model(&model.Users{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"point":      point,
			"updated_at": time.Now(),
		})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return helper.ResponseError(404, constant.ERROR_DATA_ID)
	}

	return nil
}
