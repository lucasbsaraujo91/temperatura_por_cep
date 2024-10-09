package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"temperatura_por_cep/internal/infra/api/api"
	"temperatura_por_cep/internal/infra/api/entity"
)

func FetchAddress(cep string, fetcher api.AddressFetcher) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	brasilAPIChan := make(chan entity.BrasilAPIAddress)
	viacepChan := make(chan entity.ViaCEPAddress)
	errors := make(chan error, 2)

	var once sync.Once
	var result interface{}
	var resultErr error

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

	done := make(chan struct{})

	go func() {
		select {
		case address := <-brasilAPIChan:
			once.Do(func() {
				result = address
				resultErr = nil
				fmt.Printf("Address from BrasilAPI: %+v\n", address)
				close(done)
			})
		case address := <-viacepChan:
			once.Do(func() {
				result = address
				resultErr = nil
				fmt.Printf("Address from ViaCEP: %+v\n", address)
				close(done)
			})
		case err := <-errors:
			once.Do(func() {
				result = nil
				resultErr = err
				close(done)
			})
		case <-ctx.Done():
			once.Do(func() {
				result = nil
				resultErr = fmt.Errorf("timeout while fetching address")
				close(done)
			})
		}
	}()

	<-done

	return result, resultErr
}
