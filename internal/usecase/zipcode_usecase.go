package usecase

import (
	"temperatura_por_cep/internal/entity"
	"temperatura_por_cep/internal/infra/api_busca_cep/api"
	"temperatura_por_cep/internal/infra/api_busca_cep/service"
	"temperatura_por_cep/internal/utils"
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

	err := utils.IsValid(input.ZipCode)
	if err != nil {
		return nil, err
	}

	// Chama o serviço para buscar os dados do endereço
	addressData, err := service.FetchAddress(input.ZipCode, u.Fetcher)
	if err != nil {
		return nil, err
	}

	// Cria a entidade Address
	address, err := entity.NewConsultZipCode(
		addressData.ZipCode,
		addressData.Street,
		addressData.Neighborhood,
		addressData.City,
		addressData.State,
	)
	if err != nil {
		return nil, err
	}

	// Preenche o DTO de saída
	output := &GetAddressOutputDTO{
		ZipCode:      address.ZipCode,
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
	}

	return output, nil
}
