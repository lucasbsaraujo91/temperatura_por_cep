package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"temperatura_por_cep/internal/infra/api_busca_cep/entity"
)

// Remove the redeclaration of DefaultAddressFetcher
// type DefaultAddressFetcher struct{}

func (f *DefaultAddressFetcher) FetchAddressFromViaCEP(cep string) (entity.ViaCEPAddress, error) {
	//time.Sleep(2 * time.Second)
	var address entity.ViaCEPAddress
	resp, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return address, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return address, fmt.Errorf("failed to fetch address from ViaCEP: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return address, err
	}
	return address, nil
}
