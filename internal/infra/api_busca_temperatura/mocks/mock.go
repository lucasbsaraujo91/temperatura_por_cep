package mocks

import (
	"temperatura_por_cep/internal/infra/api_busca_temperatura/entity"

	"github.com/stretchr/testify/mock"
)

// MockWeatherFetcher é um mock de WeatherFetcher.
type MockWeatherFetcher struct {
	mock.Mock
}

// FetchWeather simula a chamada à API de clima.
func (m *MockWeatherFetcher) FetchWeather(city string) (entity.WeatherResponse, error) {
	args := m.Called(city)
	return args.Get(0).(entity.WeatherResponse), args.Error(1)
}
