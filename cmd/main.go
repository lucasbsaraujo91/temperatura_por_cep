package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"temperatura_por_cep/internal/infra/api_busca_cep/api"
	teste "temperatura_por_cep/internal/infra/api_busca_temperatura/api"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/service"
	"temperatura_por_cep/internal/usecase"
)

func AddressHandler(fetcher api.AddressFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cep := chi.URLParam(r, "cep")

		addressUseCase := usecase.AddressUseCase{Fetcher: fetcher}
		address, err := addressUseCase.GetAddressByZipCode(usecase.GetAddressInputDTO{ZipCode: cep})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching address: %v", err), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(address)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling address: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// ew function to fetch weather data
func WeatherHandler(service *service.WeatherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		city := chi.URLParam(r, "city")
		country := chi.URLParam(r, "country")

		city = strings.ReplaceAll(city, " ", "+")
		country = strings.ReplaceAll(country, " ", "+")

		weatherData, err := service.FetchWeatherByCity(city, country) // Chama o método do serviço
		if err != nil {
			http.Error(w, fmt.Sprintf("Errorrr fetching weather: %v", err), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(weatherData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling weather data: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func main() {
	r := chi.NewRouter()

	fetcher := &api.DefaultAddressFetcher{}

	apiKey := "4457481f588941748b8232000240910"       // Insira sua chave da API
	weatherFetcher := teste.NewWeatherFetcher(apiKey) // Passa a chave para o fetcher
	weatherService := service.NewWeatherService(apiKey, weatherFetcher)

	// Criar a instância do WeatherUseCase
	addressUseCase := usecase.NewAddressUseCase(fetcher)
	weatherUseCase := &usecase.WeatherUseCase{
		AddressUseCase: addressUseCase,
		WeatherService: weatherService,
	}

	// Rota para buscar o clima pelo CEP
	r.Get("/weather/zipcode/{zipCode}", func(w http.ResponseWriter, r *http.Request) {
		zipCode := chi.URLParam(r, "zipCode")
		weather, err := weatherUseCase.GetWeatherByZipCode(zipCode)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching weather: %v", err), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(weather)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling weather data: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	r.Get("/address/{cep}", AddressHandler(fetcher))

	http.ListenAndServe(":8060", r)
}
