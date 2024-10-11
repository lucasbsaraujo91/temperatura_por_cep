package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"temperatura_por_cep/internal/infra/api_busca_cep/api"
	apiTemperatura "temperatura_por_cep/internal/infra/api_busca_temperatura/api"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/service"
	"temperatura_por_cep/internal/usecase"
)

func AddressHandler(fetcher api.AddressFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cep := chi.URLParam(r, "cep")
		handleAddressRequest(w, cep, fetcher)
	}
}

func handleAddressRequest(w http.ResponseWriter, cep string, fetcher api.AddressFetcher) {
	addressUseCase := usecase.NewAddressUseCase(fetcher)
	address, err := addressUseCase.GetAddressByZipCode(usecase.GetAddressInputDTO{ZipCode: cep})
	if err != nil {
		switch err.Error() {
		case "invalid zipcode":
			http.Error(w, err.Error(), http.StatusUnprocessableEntity) // 422
			return
		case "can not find zipcode":
			http.Error(w, err.Error(), http.StatusNotFound) // 404
			return
		default:
			http.Error(w, err.Error(), http.StatusNotFound) // 404
			return
		}
	}

	sendJSONResponse(w, address)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	r := chi.NewRouter()

	fetcher := &api.DefaultAddressFetcher{}
	apiKey := "4457481f588941748b8232000240910" // Insira sua chave da API
	weatherFetcher := apiTemperatura.NewWeatherFetcher(apiKey)
	weatherService := service.NewWeatherService(apiKey, weatherFetcher)

	addressUseCase := usecase.NewAddressUseCase(fetcher)
	// Rota para buscar o clima pelo CEP
	weatherUseCase := &usecase.WeatherUseCase{
		AddressUseCase: addressUseCase,
		WeatherService: weatherService,
	}
	r.Get("/weather/zipcode/{zipCode}", func(w http.ResponseWriter, r *http.Request) {
		zipCode := chi.URLParam(r, "zipCode")
		handleWeatherByZipCode(w, zipCode, weatherUseCase)
	})

	r.Get("/address/{cep}", AddressHandler(fetcher))

	http.ListenAndServe(":8060", r)
}

func handleWeatherByZipCode(w http.ResponseWriter, zipCode string, weatherUseCase *usecase.WeatherUseCase) {
	weather, err := weatherUseCase.GetWeatherByZipCode(zipCode)

	if err != nil {
		// Verifica se o erro é um erro de CEP inválido
		if err.Error() == "invalid zipcode" {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		// Verifica se o erro é de CEP não encontrado
		if err.Error() == "failed to fetch address from BrasilAPI: 404 Not Found" ||
			err.Error() == "invalid zipcode: must be 8 digits" {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		// Para outros erros, retorna erro interno
		http.Error(w, fmt.Sprintf("Error fetching weather: %v", err), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, weather)
}
