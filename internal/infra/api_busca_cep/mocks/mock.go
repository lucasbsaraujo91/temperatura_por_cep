package mocks

import (
	"temperatura_por_cep/internal/infra/api_busca_cep/entity"

	"github.com/stretchr/testify/mock"
)

// MockAddressFetcher é uma estrutura que implementa a interface de busca de endereço
type MockAddressFetcher struct {
	mock.Mock
}

func (m *MockAddressFetcher) FetchAddressFromBrasilAPI(zipCode string) (entity.BrasilAPIAddress, error) {
	args := m.Called(zipCode)
	return args.Get(0).(entity.BrasilAPIAddress), args.Error(1)
}

func (m *MockAddressFetcher) FetchAddressFromViaCEP(zipCode string) (entity.ViaCEPAddress, error) {
	args := m.Called(zipCode)
	return args.Get(0).(entity.ViaCEPAddress), args.Error(1)
}
