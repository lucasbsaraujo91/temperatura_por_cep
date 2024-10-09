package service

import "temperatura_por_cep/internal/infra/api/entity"

type AddressFetcher interface {
	FetchFromBrasilAPI(zipCode string) (entity.BrasilAPIAddress, error)
	FetchFromViaCEP(zipCode string) (entity.ViaCEPAddress, error)
}

type AddressData struct {
	CEP          string
	Street       string
	Neighborhood string
	City         string
	State        string
}
