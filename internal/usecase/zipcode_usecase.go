package usecase

import (
	"fmt"
	"temperatura_por_cep/internal/entity"
	"temperatura_por_cep/internal/infra/api/api"
	"temperatura_por_cep/internal/infra/api/service"
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
	fetcher api.AddressFetcher
}

func NewAddressUseCase(fetcher api.AddressFetcher) *AddressUseCase {
	return &AddressUseCase{fetcher: fetcher}
}

func (u *AddressUseCase) GetAddressByZipCode(input GetAddressInputDTO) (*entity.Address, error) {
	addressData, err := service.FetchAddress(input.ZipCode, u.fetcher)

	if err != nil {
		return nil, err
	}

	// Imprimindo o conteúdo do addressData
	fmt.Printf("Address: %+v\n", addressData)

	// Criando um novo Address com os dados obtidos
	newAddress, err := entity.NewConsultZipCode(
		addressData.ZipCode,
		addressData.Street,
		addressData.Neighborhood,
		addressData.City,
		addressData.State,
	)
	if err != nil {
		return nil, err
	}

	return newAddress, nil
}
