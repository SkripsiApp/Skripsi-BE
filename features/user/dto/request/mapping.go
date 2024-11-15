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

func UserRequestOTPToUserCore(data UserSendOTP) entity.UsersCore {
	return entity.UsersCore{
		Email: data.Email,
	}
}

func UserRequestVerifyOTPToUserCore(data UserVerifyOTP) entity.UsersCore {
	return entity.UsersCore{
		Email: data.Email,
		Otp:   data.Otp,
	}
}

func UserRequestNewPasswordToUserCore(data UserNewPassword) entity.UsersCore {
	return entity.UsersCore{
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}
