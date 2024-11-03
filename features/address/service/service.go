package service

import (
	"regexp"
	"skripsi/features/address/entity"
	"skripsi/features/address/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
)

type addressService struct {
	addressRepository interfaces.AddressRepositoryInterface
}

func NewAddressService(addressRepository interfaces.AddressRepositoryInterface) interfaces.AddressServiceInterface {
	return &addressService{
		addressRepository: addressRepository,
	}
}

// Create implements interfaces.AddressServiceInterface.
func (a *addressService) Create(data entity.AddressCore) (entity.AddressCore, error) {
	if data.UserId == "" {
		return entity.AddressCore{}, helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	if data.Name == "" || data.Address == "" || data.City == "" || data.Subdistric == "" || data.ZipCode == "" || data.Phone == "" {
		return entity.AddressCore{}, helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	phoneRegex := regexp.MustCompile(`^[0-9]+$`)
	if len(data.Phone) < 10 || len(data.Phone) > 15 || !phoneRegex.MatchString(data.Phone) {
		return entity.AddressCore{}, helper.ResponseError(400, "format nomor telepon tidak valid")
	}

	zipCodeRegex := regexp.MustCompile(`^[0-9]+$`)
	if len(data.ZipCode) != 5 || !zipCodeRegex.MatchString(data.ZipCode) {
		return entity.AddressCore{}, helper.ResponseError(400, "format kode pos tidak valid")
	}

	response, err := a.addressRepository.Create(data)
	if err != nil {
		return entity.AddressCore{}, err
	}

	return response, nil

}

// DeleteById implements interfaces.AddressServiceInterface.
func (a *addressService) DeleteById(id string, userId string) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	address, err := a.addressRepository.FindById(id)
	if err != nil {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	if address.UserId != userId {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	errDelete := a.addressRepository.DeleteById(id, userId)
	if errDelete != nil {
		return errDelete
	}

	return nil
}

// GetAll implements interfaces.AddressServiceInterface.
func (a *addressService) GetAll(idUser string) ([]entity.AddressCore, error) {
	if idUser == "" {
		return nil, helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	data, err := a.addressRepository.GetAll(idUser)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetById implements interfaces.AddressServiceInterface.
func (a *addressService) GetById(id string, userId string) (entity.AddressCore, error) {
	if id == "" {
		return entity.AddressCore{}, helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	data, err := a.addressRepository.GetById(id, userId)
	if err != nil {
		return entity.AddressCore{}, err
	}

	return data, nil
}

// UpdateById implements interfaces.AddressServiceInterface.
func (a *addressService) UpdateById(id string, data entity.AddressCore) error {
	panic("unimplemented")
}
