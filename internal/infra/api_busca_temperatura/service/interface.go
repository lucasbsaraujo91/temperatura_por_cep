package service

import (
	"temperatura_por_cep/internal/infra/api_busca_temperatura/entity"
)

// WeatherFetcher define a interface para buscar o clima
type WeatherFetcher interface {
	FetchWeather(city, country string) (entity.FullWeather, error)
}
