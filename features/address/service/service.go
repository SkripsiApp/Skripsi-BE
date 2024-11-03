package service

import (
	"regexp"
	"skripsi/features/address/entity"
	"skripsi/features/address/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"
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
func (a *addressService) DeleteById(id string) error {
	panic("unimplemented")
}

// GetAll implements interfaces.AddressServiceInterface.
func (a *addressService) GetAll(search string, page int, limit int) ([]entity.AddressCore, pagination.PageInfo, int, error) {
	panic("unimplemented")
}

// GetById implements interfaces.AddressServiceInterface.
func (a *addressService) GetById(id string) (entity.AddressCore, error) {
	panic("unimplemented")
}

// UpdateById implements interfaces.AddressServiceInterface.
func (a *addressService) UpdateById(id string, data entity.AddressCore) error {
	panic("unimplemented")
}
