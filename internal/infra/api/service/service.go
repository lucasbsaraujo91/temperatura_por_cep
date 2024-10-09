package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"temperatura_por_cep/internal/infra/api/api"
	"temperatura_por_cep/internal/infra/api/entity"
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

	var once sync.Once
	var result AddressDTO
	var resultErr error

	go func() {
		address, err := fetcher.FetchAddressFromBrasilAPI(cep)
		if err != nil {
			errors <- err
			return
		}
		// Converte para DTO
		result = AddressDTO{
			ZipCode:      address.CEP,
			Street:       address.Street,
			Neighborhood: address.Neighborhood,
			City:         address.City,
			State:        address.State,
		}
		brasilAPIChan <- address
	}()

	go func() {
		address, err := fetcher.FetchAddressFromViaCEP(cep)
		if err != nil {
			errors <- err
			return
		}
		// Converte para DTO
		result = AddressDTO{
			ZipCode:      address.CEP,
			Street:       address.Logradouro,
			Neighborhood: address.Bairro,
			City:         address.Localidade,
			State:        address.UF,
		}
		viacepChan <- address
	}()

	done := make(chan struct{})

	go func() {
		select {
		case <-brasilAPIChan:
			once.Do(func() {
				resultErr = nil
				fmt.Printf("Address from BrasilAPI: %+v\n", result)
				close(done)
			})
		case <-viacepChan:
			once.Do(func() {
				resultErr = nil
				fmt.Printf("Address from ViaCEP: %+v\n", result)
				close(done)
			})
		case err := <-errors:
			once.Do(func() {
				resultErr = err
				close(done)
			})
		case <-ctx.Done():
			once.Do(func() {
				resultErr = fmt.Errorf("timeout while fetching address")
				close(done)
			})
		}
	}()

	<-done

	return result, resultErr
}
