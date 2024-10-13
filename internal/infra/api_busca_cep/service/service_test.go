package service_test

import (
	"errors"
	"testing"
	"time"

	"temperatura_por_cep/internal/infra/api_busca_cep/entity"
	"temperatura_por_cep/internal/infra/api_busca_cep/mocks"
	"temperatura_por_cep/internal/infra/api_busca_cep/service"
)

func TestFetchAddress_Success(t *testing.T) {
	// Mock do AddressFetcher
	fetcher := &mocks.MockAddressFetcher{}

	// Dados de retorno simulados
	fetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{
		CEP:          "12345678",
		Street:       "Rua Exemplo",
		Neighborhood: "Bairro Exemplo",
		City:         "São Paulo",
		State:        "SP",
	}, nil)

	fetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{
		CEP:        "12345678",
		Logradouro: "Rua Exemplo",
		Bairro:     "Bairro Exemplo",
		Localidade: "São Paulo",
		UF:         "SP",
	}, nil)

	// Chama a função que queremos testar
	addressDTO, err := service.FetchAddress("12345678", fetcher)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verifica os dados retornados
	if addressDTO.ZipCode != "12345678" {
		t.Errorf("Expected ZipCode to be '12345678', got %s", addressDTO.ZipCode)
	}
	if addressDTO.Street != "Rua Exemplo" {
		t.Errorf("Expected Street to be 'Rua Exemplo', got %s", addressDTO.Street)
	}
	if addressDTO.Neighborhood != "Bairro Exemplo" {
		t.Errorf("Expected Neighborhood to be 'Bairro Exemplo', got %s", addressDTO.Neighborhood)
	}
	if addressDTO.City != "São Paulo" {
		t.Errorf("Expected City to be 'São Paulo', got %s", addressDTO.City)
	}
	if addressDTO.State != "SP" {
		t.Errorf("Expected State to be 'SP', got %s", addressDTO.State)
	}
}

func TestFetchAddress_Error(t *testing.T) {
	// Mock do AddressFetcher
	fetcher := &mocks.MockAddressFetcher{}

	// Simulando erro da API
	fetcher.On("FetchAddressFromBrasilAPI", "invalid").Return(entity.BrasilAPIAddress{}, errors.New("invalid zipcode"))

	// Chama a função que queremos testar
	addressDTO, err := service.FetchAddress("invalid", fetcher)

	// Verifica se o erro foi retornado
	if err == nil {
		t.Fatalf("Expected error, got none")
	}

	if addressDTO.ZipCode != "" {
		t.Errorf("Expected no DTO, got %+v", addressDTO)
	}
}

func TestFetchAddress_Timeout(t *testing.T) {
	fetcher := new(mocks.MockAddressFetcher)

	// Simula que o método de BrasilAPI irá retornar um erro após um delay
	fetcher.On("FetchAddressFromBrasilAPI", "12345678").Return(entity.BrasilAPIAddress{}, nil).After(2 * time.Second)

	// Simula que o método de ViaCEP irá retornar um erro após um delay
	fetcher.On("FetchAddressFromViaCEP", "12345678").Return(entity.ViaCEPAddress{}, nil).After(2 * time.Second)

	_, err := service.FetchAddress("12345678", fetcher)

	if err == nil || err.Error() != "timeout while fetching address" {
		t.Errorf("Expected timeout error, got %v", err)
	}

	// Verifica se os métodos foram chamados corretamente
	fetcher.AssertExpectations(t)
}
