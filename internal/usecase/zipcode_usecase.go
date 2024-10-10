package usecase

import (
	"temperatura_por_cep/internal/infra/api_busca_cep/api"
	"temperatura_por_cep/internal/infra/api_busca_cep/service"
)

type GetAddressInputDTO struct {
	ZipCode string `json:"zipcode"`
}

type GetAddressOutputDTO struct {
	ZipCode      string `json:"zipcode"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type AddressUseCase struct {
	Fetcher api.AddressFetcher
}

func NewAddressUseCase(fetcher api.AddressFetcher) *AddressUseCase {
	return &AddressUseCase{Fetcher: fetcher}
}

func (u *AddressUseCase) GetAddressByZipCode(input GetAddressInputDTO) (*GetAddressOutputDTO, error) {
	addressData, err := service.FetchAddress(input.ZipCode, u.Fetcher)

	if err != nil {
		return nil, err
	}

	output := &GetAddressOutputDTO{
		ZipCode:      addressData.ZipCode,
		Street:       addressData.Street,
		Neighborhood: addressData.Neighborhood,
		City:         addressData.City,
		State:        addressData.State,
	}

	return output, nil
}
