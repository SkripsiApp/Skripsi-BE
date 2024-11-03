package mapping

import (
	"skripsi/features/address/entity"
	"skripsi/features/address/model"
)

// Mapping Core to Model
func AddressCoreToAddressModel(data entity.AddressCore) model.Address {
	addressModel := model.Address{
		Id:         data.Id,
		UserId:     data.UserId,
		Name:       data.Name,
		Address:    data.Address,
		City:       data.City,
		Subdistric: data.Subdistric,
		ZipCode:    data.ZipCode,
		Phone:      data.Phone,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeleteAt:   data.DeleteAt,
	}
	return addressModel
}

func ListAddressCoreToAddressModel(data []entity.AddressCore) []model.Address {
	listAddressModel := []model.Address{}
	for _, addressCore := range data {
		addressModel := AddressCoreToAddressModel(addressCore)
		listAddressModel = append(listAddressModel, addressModel)
	}
	return listAddressModel
}

// Mapping Model to Core
func AddressModelToAddressCore(data model.Address) entity.AddressCore {
	addressCore := entity.AddressCore{
		Id:         data.Id,
		UserId:     data.UserId,
		Name:       data.Name,
		Address:    data.Address,
		City:       data.City,
		Subdistric: data.Subdistric,
		ZipCode:    data.ZipCode,
		Phone:      data.Phone,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeleteAt:   data.DeleteAt,
	}
	return addressCore
}

func ListAddressModelToAddressCore(data []model.Address) []entity.AddressCore {
	listAddressCore := []entity.AddressCore{}
	for _, addressModel := range data {
		addressCore := AddressModelToAddressCore(addressModel)
		listAddressCore = append(listAddressCore, addressCore)
	}
	return listAddressCore
}
