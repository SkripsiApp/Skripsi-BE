package request

import "skripsi/features/user/entity"

func UserRegisterToUserCore(data UserRegister) entity.UsersCore {
	return entity.UsersCore{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}
