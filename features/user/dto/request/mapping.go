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

func UserUpdateToUserCore(data UserUpdate) entity.UsersCore {
	return entity.UsersCore{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
	}
}

func UserLoginToUserCore(data UserLogin) entity.UsersCore {
	return entity.UsersCore{
		Email:    data.Email,
		Password: data.Password,
	}
}
