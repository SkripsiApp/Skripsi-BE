package response

import "skripsi/features/admin/entity"

func AdminCoreToLoginResponse(data entity.AdminCore, token string) LoginResponse {
	return LoginResponse{
		Id:    data.Id,
		Email: data.Email,
		Token: token,
	}
}
