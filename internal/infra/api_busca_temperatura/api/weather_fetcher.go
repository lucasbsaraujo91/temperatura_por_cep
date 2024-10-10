package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/entity"
)

type WeatherFetcherImpl struct{}

func (o *WeatherFetcherImpl) FetchWeather(city, country, apiKey string) (entity.Current, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s,%s&lang=pt_br", apiKey, city, country)

	// Faz a requisição HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return entity.Current{}, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		return entity.Current{}, fmt.Errorf("erro: status code %d", resp.StatusCode)
	}

	// Decodifica o corpo da resposta para a struct WeatherResponse
	var weather entity.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return entity.Current{}, fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	return weather.Current, nil
}
