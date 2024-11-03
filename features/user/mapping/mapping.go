package mapping

import (
	"skripsi/features/user/entity"
	"skripsi/features/user/model"
)

// Mapping Core to Model
func UserCoreToUserModel(data entity.UsersCore) model.Users {
	userModel := model.Users{
		Id:         data.Id,
		Name:       data.Name,
		Username:   data.Username,
		Email:      data.Email,
		Password:   data.Password,
		Point:      data.Point,
		Role:       data.Role,
		Otp:        data.Otp,
		OtpExpired: data.OtpExpired,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeleteAt:   data.DeleteAt,
	}
	return userModel
}

func ListUserCoreToUserModel(data []entity.UsersCore) []model.Users {
	listUserModel := []model.Users{}
	for _, userCore := range data {
		userModel := UserCoreToUserModel(userCore)
		listUserModel = append(listUserModel, userModel)
	}
	return listUserModel
}

// Mapping Model to Core
func UserModelToUserCore(data model.Users) entity.UsersCore {
	userCore := entity.UsersCore{
		Id:              data.Id,
		Name:            data.Name,
		Username:        data.Username,
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.Password,
		Point:           data.Point,
		Role:            data.Role,
		Otp:             data.Otp,
		OtpExpired:      data.OtpExpired,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
		DeleteAt:        data.DeleteAt,
	}
	return userCore
}

func ListUserModelToUserCore(data []model.Users) []entity.UsersCore {
	listUserCore := []entity.UsersCore{}
	for _, userModel := range data {
		userCore := UserModelToUserCore(userModel)
		listUserCore = append(listUserCore, userCore)
	}
	return listUserCore
}


