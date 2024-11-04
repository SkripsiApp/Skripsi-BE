package request

import "skripsi/features/admin/entity"

func AdminRegisterToAdminCore(data AdminRegister) entity.AdminCore {
	return entity.AdminCore{
		Email:    data.Email,
		Password: data.Password,
	}
}

func AdminLoginToAdminCore(data AdminLogin) entity.AdminCore {
	return entity.AdminCore{
		Email:    data.Email,
		Password: data.Password,
	}
}