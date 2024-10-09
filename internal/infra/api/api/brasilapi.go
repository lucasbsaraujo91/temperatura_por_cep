package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"temperatura_por_cep/internal/infra/api/entity"
)

type DefaultAddressFetcher struct{}

func (f *DefaultAddressFetcher) FetchAddressFromBrasilAPI(cep string) (entity.BrasilAPIAddress, error) {
	//time.Sleep(2 * time.Second)
	var address entity.BrasilAPIAddress
	resp, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
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
