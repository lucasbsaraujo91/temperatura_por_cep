package usecase

import (
	"fmt"
	"testing"

	"temperatura_por_cep/internal/infra/api_busca_cep/entity"
	"temperatura_por_cep/internal/infra/api_busca_cep/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressByZipCode_Success(t *testing.T) {
	mockFetcher := new(mocks.MockAddressFetcher)

	// Simule o retorno da API BrasilAPI
	mockFetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{
		CEP:          "12345678",
		Street:       "Rua Exemplo",
		Neighborhood: "Bairro Exemplo",
		City:         "São Paulo",
		State:        "SP",
	}, nil)

	// Simule o retorno da API ViaCEP se for chamado
	mockFetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{
		CEP:        "12345678",
		Logradouro: "Rua Exemplo",
		Bairro:     "Bairro Exemplo",
		Localidade: "São Paulo",
		UF:         "SP",
	}, nil)

	// Crie o caso de uso com o mock
	addressUseCase := NewAddressUseCase(mockFetcher)

	// Chame o método a ser testado
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "12345678"})

	// Verifique os resultados
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "12345678", output.ZipCode)
	assert.Equal(t, "Rua Exemplo", output.Street)
	assert.Equal(t, "Bairro Exemplo", output.Neighborhood)
	assert.Equal(t, "São Paulo", output.City)
	assert.Equal(t, "SP", output.State)

	// Verifique se a expectativa do mock foi atendida
	mockFetcher.AssertExpectations(t)
}

func TestGetAddressByZipCode_InvalidZipCode(t *testing.T) {
	mockFetcher := new(mocks.MockAddressFetcher)

	addressUseCase := NewAddressUseCase(mockFetcher)

	// Chame o método com um CEP inválido
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "invalid"})

	// Verifique os resultados
	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestGetAddressByZipCode_NotFound(t *testing.T) {
	mockFetcher := new(mocks.MockAddressFetcher)

	// Simula o retorno de erro quando o CEP não é encontrado
	mockFetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{}, fmt.Errorf("not found"))
	mockFetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{}, fmt.Errorf("not found"))

	// Crie o caso de uso com o mock
	addressUseCase := NewAddressUseCase(mockFetcher)

	// Chame o método a ser testado
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "12345678"})

	// Verifique os resultados
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "not found", err.Error())

	// Verifique se as expectativas do mock foram atendidas
	mockFetcher.AssertExpectations(t)
}
