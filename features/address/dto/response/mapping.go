package response

import "skripsi/features/address/entity"

func AddressCoreToAddressResponse(data entity.AddressCore) AddressResponse {
	return AddressResponse{
		Id:         data.Id,
		Name:       data.Name,
		Address:    data.Address,
		City:       data.City,
		Subdistric: data.Subdistric,
		ZipCode:    data.ZipCode,
		Phone:      data.Phone,
	}
}

func ListAddressCoreToAddressResponse(data []entity.AddressCore) []AddressResponse {
	list := []AddressResponse{}
	for _, value := range data {
		result := AddressCoreToAddressResponse(value)
		list = append(list, result)
	}
	return list
}
