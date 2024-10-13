package usecase_test

import (
	"errors"
	"temperatura_por_cep/internal/infra/api_busca_cep/api"
	"temperatura_por_cep/internal/usecase"
	"testing"
)

// Mock para o AddressFetcher
type MockAddressFetcher struct{}

func (m *MockAddressFetcher) FetchAddressByZipCode(zipCode string) (api.Address, error) {
	if zipCode == "12345678" {
		return api.Address{
			ZipCode: "12345678",
			City:    "São Paulo",
		}, nil
	}
	return api.Address{}, errors.New("invalid zipcode")
}

func TestGetAddressByZipCode_ValidZipCode(t *testing.T) {
	fetcher := &MockAddressFetcher{}
	addressUseCase := usecase.NewAddressUseCase(fetcher)

	input := usecase.GetAddressInputDTO{
		ZipCode: "12345678",
	}
	address, err := addressUseCase.GetAddressByZipCode(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if address.ZipCode != "12345678" {
		t.Errorf("expected ZipCode to be '12345678', got '%s'", address.ZipCode)
	}
	if address.City != "São Paulo" {
		t.Errorf("expected City to be 'São Paulo', got '%s'", address.City)
	}
}

func TestGetAddressByZipCode_InvalidZipCode(t *testing.T) {
	fetcher := &MockAddressFetcher{}
	addressUseCase := usecase.NewAddressUseCase(fetcher)

	input := usecase.GetAddressInputDTO{
		ZipCode: "00000000",
	}
	_, err := addressUseCase.GetAddressByZipCode(input)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "invalid zipcode" {
		t.Errorf("expected error message 'invalid zipcode', got '%s'", err.Error())
	}
}
