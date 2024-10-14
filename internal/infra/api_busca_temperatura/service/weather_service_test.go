package service_test

import (
	"errors"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/entity"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/service"
	"testing"
)

// Mock da interface WeatherFetcher
type mockWeatherFetcher struct {
	mockResponse entity.FullWeather
	mockError    error
}

// Implementação do método FetchWeather no mock
func (m *mockWeatherFetcher) FetchWeather(city, country string) (entity.FullWeather, error) {
	return m.mockResponse, m.mockError
}

func TestWeatherService_FetchWeatherByCity_Success(t *testing.T) {
	// Dados simulados de retorno da API
	mockResponse := entity.FullWeather{
		Current: entity.Current{
			TempC: 22,
			Condition: entity.Condition{
				Text: "Sunny",
			},
		},
	}

	// Criar o mock do WeatherFetcher
	mockFetcher := &mockWeatherFetcher{
		mockResponse: mockResponse,
		mockError:    nil,
	}

	// Instanciar o serviço com o mock
	weatherService := service.NewWeatherService("fakeApiKey", mockFetcher)

	// Chamar a função que queremos testar
	result, err := weatherService.FetchWeatherByCity("São Paulo", "BR")

	// Verificar se não houve erro
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verificar se o resultado é o esperado
	if result.Current.TempC != 22 {
		t.Errorf("Expected temperature 22, got %v", result.Current.TempC)
	}

	if result.Current.Condition.Text != "Sunny" {
		t.Errorf("Expected condition 'Sunny', got %v", result.Current.Condition.Text)
	}
}

func TestWeatherService_FetchWeatherByCity_Error(t *testing.T) {
	// Criar o mock para retornar um erro
	mockFetcher := &mockWeatherFetcher{
		mockError: errors.New("failed to fetch weather"),
	}

	weatherService := service.NewWeatherService("fakeApiKey", mockFetcher)

	// Chamar a função que queremos testar
	_, err := weatherService.FetchWeatherByCity("São Paulo", "BR")

	// Verificar se o erro foi retornado corretamente
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	if err.Error() != "failed to fetch weather" {
		t.Errorf("Expected error 'failed to fetch weather', got %v", err)
	}
}
