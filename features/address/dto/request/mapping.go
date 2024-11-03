package request

import "skripsi/features/address/entity"

func AddressRequestToAddressCore(data AddressRequest) entity.AddressCore {
	return entity.AddressCore{
		UserId:     data.UserId,
		Name:       data.Name,
		Address:    data.Address,
		City:       data.City,
		Subdistric: data.Subdistric,
		ZipCode:    data.ZipCode,
		Phone:      data.Phone,
	}
}
