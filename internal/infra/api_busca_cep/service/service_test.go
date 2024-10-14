package service

import (
	"testing"
	"time"

	"temperatura_por_cep/internal/infra/api_busca_cep/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAddressFetcher é um mock do AddressFetcher
type MockAddressFetcher struct {
	mock.Mock
}

func (m *MockAddressFetcher) FetchAddressFromBrasilAPI(cep string) (entity.BrasilAPIAddress, error) {
	args := m.Called(cep)
	return args.Get(0).(entity.BrasilAPIAddress), args.Error(1)
}

func (m *MockAddressFetcher) FetchAddressFromViaCEP(cep string) (entity.ViaCEPAddress, error) {
	args := m.Called(cep)
	return args.Get(0).(entity.ViaCEPAddress), args.Error(1)
}

func TestFetchAddress_Success_BrasilAPI_Faster(t *testing.T) {
	mockFetcher := new(MockAddressFetcher)

	// Expectativa para BrasilAPI
	mockFetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{
		Cep:          "12345678",
		Street:       "Rua Exemplo",
		Neighborhood: "Bairro Exemplo",
		City:         "São Paulo",
		State:        "SP",
	}, nil)

	// Expectativa para ViaCEP (retorno mais lento)
	mockFetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{
		Cep:        "12345678",
		Logradouro: "Rua Exemplo",
		Bairro:     "Bairro Exemplo",
		Localidade: "São Paulo",
		UF:         "SP",
	}, nil).After(200 * time.Millisecond) // simula uma chamada mais lenta

	address, err := FetchAddress("12345678", mockFetcher)

	assert.NoError(t, err)

	expectedAddress := AddressDTO{
		ZipCode:      "12345678",
		Street:       "Rua Exemplo",
		Neighborhood: "Bairro Exemplo",
		City:         "São Paulo",
		State:        "SP",
	}

	assert.Equal(t, expectedAddress, address)

	// Verifica se todas as expectativas foram atendidas
	mockFetcher.AssertExpectations(t)
}
