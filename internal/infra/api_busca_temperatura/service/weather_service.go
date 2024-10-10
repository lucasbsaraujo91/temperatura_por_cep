package service

import (
	"temperatura_por_cep/internal/infra/api_busca_temperatura/entity"
)

// WeatherService define o serviço que obtém o clima
type WeatherService struct {
	APIKey  string
	Fetcher WeatherFetcher
}

// NewWeatherService cria uma nova instância de WeatherService
func NewWeatherService(apiKey string, fetcher WeatherFetcher) *WeatherService {
	return &WeatherService{
		APIKey:  apiKey,
		Fetcher: fetcher, // Referência ao fetcher passado como parâmetro
	}
}

// FetchWeatherByCity busca o clima para uma cidade e país
func (s *WeatherService) FetchWeatherByCity(city, country string) (entity.FullWeather, error) {
	return s.Fetcher.FetchWeather(city, country)
}
