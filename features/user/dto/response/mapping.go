package response

import "skripsi/features/user/entity"

func UserCoreToUserResponse(data entity.UsersCore) UserResponse {
	return UserResponse{
		Id:       data.Id,
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Point:    data.Point,
	}
}

func ListUserCoreToUserResponse(data []entity.UsersCore) []UserResponse {
	list := []UserResponse{}
	for _, value := range data {
		result := UserCoreToUserResponse(value)
		list = append(list, result)
	}
	return list
}
