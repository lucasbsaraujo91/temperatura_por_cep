package usecase

import (
	"testing"

	"temperatura_por_cep/internal/infra/api_busca_cep/entity"
	"temperatura_por_cep/internal/infra/api_busca_cep/mocks"
	"temperatura_por_cep/internal/infra/api_busca_cep/service"

	"github.com/stretchr/testify/assert"
)

// Mocks para as APIs

func TestFetchAddress_Success(t *testing.T) {
	// Mock do AddressFetcher
	fetcher := &mocks.MockAddressFetcher{}

	// Dados de retorno simulados para BrasilAPI
	// Dados de retorno simulados para BrasilAPI
	fetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{
		Cep:          "12345678", // Certifique-se de que este campo está correto
		Street:       "Rua Exemplo",
		Neighborhood: "Bairro Exemplo",
		City:         "São Paulo",
		State:        "SP",
	}, nil)

	// Dados de retorno simulados para ViaCEP
	fetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{
		Cep:         "12345678",
		Logradouro:  "Rua Exemplo",
		Bairro:      "Bairro Exemplo",
		Localidade:  "São Paulo",
		UF:          "SP",
		Complemento: "Complemento Exemplo",
		IBGE:        "123456",
		GIA:         "1234",
		DDD:         "11",
		SIAFI:       "1234",
	}, nil)

	// Chama a função que queremos testar
	addressDTO, err := service.FetchAddress("12345678", fetcher) // Chama a função correta do serviço
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verifica os dados retornados
	assert.Equal(t, "12345678", addressDTO.ZipCode)
	assert.Equal(t, "Rua Exemplo", addressDTO.Street)
	assert.Equal(t, "Bairro Exemplo", addressDTO.Neighborhood)
	assert.Equal(t, "São Paulo", addressDTO.City)
	assert.Equal(t, "SP", addressDTO.State)

	// Verifica se todas as expectativas foram atendidas
	fetcher.AssertExpectations(t)
}
