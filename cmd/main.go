package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"temperatura_por_cep/internal/infra/api/api"
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

func main() {
	r := chi.NewRouter()

	// Criar uma instância do fetcher padrão
	fetcher := &api.DefaultAddressFetcher{}

	// Definir rotas
	r.Get("/address/{cep}", AddressHandler(fetcher))

	http.ListenAndServe(":8060", r)
}
