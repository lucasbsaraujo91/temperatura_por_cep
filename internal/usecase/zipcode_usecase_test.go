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

	// Cria o caso de uso com o mock
	addressUseCase := NewAddressUseCase(mockFetcher)

	// Chame o método a ser testado
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "12345678"})

	// Verifique os resultados
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "Rua Exemplo", output.Street)

	// Verifique se as expectativas do mock foram atendidas
	mockFetcher.AssertExpectations(t)
}

func TestGetAddressByZipCode_Success2(t *testing.T) {
	mockFetcher := new(mocks.MockAddressFetcher)

	// Simulação de resposta bem-sucedida
	mockFetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{
		CEP:        "12345678",
		Logradouro: "Rua Exemplo",
		Bairro:     "Bairro Exemplo",
		Localidade: "São Paulo",
		UF:         "SP",
	}, nil)

	addressUseCase := NewAddressUseCase(mockFetcher)
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "12345678"})

	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "Rua Exemplo", output.Street)

	// Verifique as expectativas
	mockFetcher.AssertExpectations(t)
}

func TestGetAddressByZipCode_InvalidZipCode(t *testing.T) {
	mockFetcher := new(mocks.MockAddressFetcher)

	// Não espera chamadas para buscar o endereço quando o CEP é inválido
	addressUseCase := NewAddressUseCase(mockFetcher)

	// Tente buscar um endereço com um CEP inválido
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "invalid"})

	// Verifique o erro e a saída
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "invalid zipcode", err.Error())

	// Verifique se as expectativas do mock foram atendidas
	mockFetcher.AssertExpectations(t) // Isso deve passar, pois não há chamadas esperadas
}

func TestGetAddressByZipCode_NotFound(t *testing.T) {
	// Cria um novo mock para simular o AddressFetcher
	mockFetcher := new(mocks.MockAddressFetcher)

	// Simule o retorno da API BrasilAPI como um erro
	mockFetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{}, fmt.Errorf("not found"))

	// Simule o retorno da API ViaCEP como um erro
	mockFetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{}, fmt.Errorf("not found"))

	// Crie o caso de uso com o mock
	addressUseCase := NewAddressUseCase(mockFetcher)

	// Chame o método a ser testado
	output, err := addressUseCase.GetAddressByZipCode(GetAddressInputDTO{ZipCode: "12345678"})

	// Verifique os resultados
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "not found", err.Error()) // Verifique a mensagem de erro

	// Verifique se as expectativas do mock foram atendidas
	mockFetcher.AssertExpectations(t)
}
