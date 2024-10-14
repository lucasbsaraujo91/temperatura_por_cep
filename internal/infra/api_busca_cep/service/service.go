package service

import (
	"context"
	"fmt"
	"time"

	"temperatura_por_cep/internal/infra/api_busca_cep/api"
	"temperatura_por_cep/internal/infra/api_busca_cep/entity"
)

// AddressDTO representa os dados de um endereço.
type AddressDTO struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"cidade"`
	State        string `json:"estado"`
}

func FetchAddress(cep string, fetcher api.AddressFetcher) (AddressDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	brasilAPIChan := make(chan entity.BrasilAPIAddress)
	viacepChan := make(chan entity.ViaCEPAddress)
	errors := make(chan error, 2)

	var result AddressDTO
	//var resultErr error

	go func() {
		address, err := fetcher.FetchAddressFromBrasilAPI(cep)
		if err != nil {
			errors <- err
			return
		}
		brasilAPIChan <- address
	}()

	go func() {
		address, err := fetcher.FetchAddressFromViaCEP(cep)
		if err != nil {
			errors <- err
			return
		}
		viacepChan <- address
	}()

	select {
	case address := <-brasilAPIChan:
		result = AddressDTO{
			ZipCode:      address.Cep,
			Street:       address.Street,
			Neighborhood: address.Neighborhood,
			City:         address.City,
			State:        address.State,
		}
		fmt.Printf("Address from BrasilAPI: %+v\n", result)
		return result, nil
	case address := <-viacepChan:
		result = AddressDTO{
			ZipCode:      address.Cep,
			Street:       address.Logradouro,
			Neighborhood: address.Bairro,
			City:         address.Localidade,
			State:        address.UF,
		}
		fmt.Printf("Address from ViaCEP: %+v\n", result)
		return result, nil
	case err := <-errors:
		return result, err
	case <-ctx.Done():
		return result, fmt.Errorf("timeout while fetching address")
	}
}
