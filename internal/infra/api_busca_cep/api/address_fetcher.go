package api

import "temperatura_por_cep/internal/infra/api_busca_cep/entity"

type AddressFetcher interface {
	FetchAddressFromBrasilAPI(cep string) (entity.BrasilAPIAddress, error)
	FetchAddressFromViaCEP(cep string) (entity.ViaCEPAddress, error)
}
