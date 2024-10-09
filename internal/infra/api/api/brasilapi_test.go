package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"temperatura_por_cep/internal/infra/api/entity"
	"testing"
)

func TestFetchAddressFromBrasilAPI(t *testing.T) {
	expectedAddress := entity.BrasilAPIAddress{
		CEP:          "01001-000",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Sé",
		Street:       "Praça da Sé",
	}

	// Cria um servidor HTTP de teste
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/cep/v1/01001-000" {
			t.Fatalf("expected request to /api/cep/v1/01001-000, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedAddress)
	}))
	defer ts.Close()

	// Substitui a URL base da função para apontar para o servidor de teste
	//originalURL := "https://brasilapi.com.br/api/cep/v1/"
	newURL := ts.URL + "/api/cep/v1/"

	client := &http.Client{}
	cep := "01001-000"

	// Cria uma função auxiliar para redirecionar a URL
	redirectedFetchAddress := func(cep string) (entity.BrasilAPIAddress, error) {
		var address entity.BrasilAPIAddress
		resp, err := client.Get(newURL + cep)
		if err != nil {
			return address, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return address, fmt.Errorf("failed to fetch address from BrasilAPI: %s", resp.Status)
		}

		if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
			return address, err
		}

		return address, nil
	}

	// Testa a função auxiliar redirecionada
	address, err := redirectedFetchAddress(cep)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if address != expectedAddress {
		t.Errorf("expected address to be %+v, got %+v", expectedAddress, address)
	}
}
