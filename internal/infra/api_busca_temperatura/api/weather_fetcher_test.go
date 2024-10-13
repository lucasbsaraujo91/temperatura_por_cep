package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"temperatura_por_cep/internal/infra/api_busca_temperatura/api"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/entity"

	"github.com/stretchr/testify/assert"
)

func TestFetchWeather_Success(t *testing.T) {
	// Criar um servidor HTTP mock
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := entity.FullWeather{
			Location: entity.Location{
				Name:        "São Paulo",
				Region:      "SP",
				Country:     "Brasil",
				Lat:         -23.5505,
				Lon:         -46.6333,
				TzID:        "America/Sao_Paulo",
				Localtime:   "2024-10-13 19:00",
				LocaltimeEp: 1728856800,
			},
			Current: entity.Current{
				LastUpdatedEp: 1728856800,
				LastUpdated:   "2024-10-13 19:00",
				TempC:         25.5,
				TempF:         77.9,
				Condition: entity.Condition{
					Text: "Sunny",
				},
			},
		}

		// Configura o cabeçalho da resposta
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Cria um WeatherFetcher com a URL do servidor mock
	weatherFetcher := api.NewWeatherFetcher("sua_api_key")
	weatherFetcher.APIUrl = server.URL // Usando o URL do servidor de teste

	// Chama a função FetchWeather
	result, err := weatherFetcher.FetchWeather("São+Paulo", "BR")

	// Verifica os resultados
	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", result.Location.Name)
	assert.Equal(t, 25.5, result.Current.TempC)
	assert.Equal(t, "Sunny", result.Current.Condition.Text)
}

func TestFetchWeather_NetworkError(t *testing.T) {
	// Criar um servidor HTTP mock que retorna um erro específico
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Simula um erro no servidor com status 400
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"error": "Bad Request",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response) // Retorna um JSON com o erro
	}

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Cria um WeatherFetcher com a URL do servidor mock
	weatherFetcher := api.NewWeatherFetcher("sua_api_key")
	weatherFetcher.APIUrl = server.URL // Usando o URL do servidor de teste

	// Chama a função FetchWeather
	_, err := weatherFetcher.FetchWeather("São+Paulo", "BR")

	// Verifica se o erro foi retornado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro: status code 400")
}

func TestFetchWeather_DecodingError(t *testing.T) {
	// Criar um servidor HTTP mock que retorna um conteúdo inválido
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Retorna um conteúdo que não é um JSON válido
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "This is not a valid JSON") // Retorna uma string simples
	}

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Cria um WeatherFetcher com a URL do servidor mock
	weatherFetcher := api.NewWeatherFetcher("sua_api_key")
	weatherFetcher.APIUrl = server.URL // Usando o URL do servidor de teste

	// Chama a função FetchWeather
	_, err := weatherFetcher.FetchWeather("São+Paulo", "BR")

	// Verifica se o erro foi retornado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao decodificar resposta")
}

func TestFetchWeather_EmptyResponse(t *testing.T) {
	// Criar um servidor HTTP mock que retorna uma resposta vazia
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Retorna status 200
		// Não envia conteúdo no corpo da resposta
	}

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Cria um WeatherFetcher com a URL do servidor mock
	weatherFetcher := api.NewWeatherFetcher("sua_api_key")
	weatherFetcher.APIUrl = server.URL // Usando o URL do servidor de teste

	// Chama a função FetchWeather
	_, err := weatherFetcher.FetchWeather("São+Paulo", "BR")

	// Verifica se o erro foi retornado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao decodificar resposta") // Verifica se a mensagem de erro é a esperada
}

func TestFetchWeather_InvalidLocation(t *testing.T) {
	// Criar um servidor HTTP mock que retorna um erro específico para localização inválida
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Simula um erro no servidor com status 404 para localização não encontrada
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{
			"error": "Location not found",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response) // Retorna um JSON com o erro
	}

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Cria um WeatherFetcher com a URL do servidor mock
	weatherFetcher := api.NewWeatherFetcher("sua_api_key")
	weatherFetcher.APIUrl = server.URL // Usando o URL do servidor de teste

	// Chama a função FetchWeather com uma localização inválida
	_, err := weatherFetcher.FetchWeather("Localização Inexistente", "BR")

	// Verifica se o erro foi retornado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro: status code 400")
}
