package usecase_test

import (
	"fmt"
	"temperatura_por_cep/internal/infra/api_busca_cep/entity"
	ab "temperatura_por_cep/internal/infra/api_busca_temperatura/entity"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/service"
	"temperatura_por_cep/internal/usecase"
	"testing"
)

// Mock para AddressFetcher
type mockAddressFetcher struct {
	mockResponse *entity.BrasilAPIAddress
	mockError    error
}

// Mock para WeatherFetcher
type mockWeatherFetcher struct {
	mockResponse ab.FullWeather
	mockError    error
}

func (m *mockWeatherFetcher) FetchWeather(city, state string) (ab.FullWeather, error) {
	if m.mockError != nil {
		return ab.FullWeather{}, m.mockError
	}
	return m.mockResponse, nil
}

func (m *mockAddressFetcher) FetchAddressFromViaCEP(zipCode string) (entity.ViaCEPAddress, error) {
	return entity.ViaCEPAddress{
		Logradouro: m.mockResponse.Street,
		Localidade: m.mockResponse.City,
		UF:         m.mockResponse.State,
		Cep:        m.mockResponse.Cep,
	}, nil
}

func (m *mockAddressFetcher) FetchAddressFromBrasilAPI(zipCode string) (entity.BrasilAPIAddress, error) {
	if m.mockError != nil {
		return entity.BrasilAPIAddress{}, m.mockError
	}
	return *m.mockResponse, nil
}

func TestWeatherUseCase_GetWeatherByZipCode_Success(t *testing.T) {
	mockAddressResponse := &entity.BrasilAPIAddress{
		Street: "Rua Fictícia",
		City:   "São Paulo",
		State:  "SP",
		Cep:    "06765000",
	}

	mockWeatherResponse := ab.FullWeather{
		Current: ab.Current{
			TempC: 25,
		},
	}

	// Configurando mocks
	mockAddressFetcher := &mockAddressFetcher{
		mockResponse: mockAddressResponse,
		mockError:    nil,
	}

	mockWeatherFetcher := &mockWeatherFetcher{
		mockResponse: mockWeatherResponse,
		mockError:    nil,
	}

	// Instanciando o WeatherService com o mockWeatherFetcher
	weatherService := service.NewWeatherService("fake_api_key", mockWeatherFetcher)

	// Instanciando o WeatherUseCase com os mocks
	weatherUseCase := &usecase.WeatherUseCase{
		AddressUseCase: &usecase.AddressUseCase{Fetcher: mockAddressFetcher},
		WeatherService: weatherService,
	}

	// Executando o método de teste
	result, err := weatherUseCase.GetWeatherByZipCode("06765000")
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.TempC != 25 {
		t.Errorf("Expected temperature 25, got %v", result.TempC)
	}
}
