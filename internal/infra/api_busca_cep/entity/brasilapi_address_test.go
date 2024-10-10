package entity

import (
	"encoding/json"
	"testing"
)

func TestBrasilAddressCreation(t *testing.T) {

	address := BrasilAPIAddress{
		CEP:          "12345678",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Vila Mariana",
		Street:       "Rua Domingos de Morais",
	}

	if address.CEP != "12345678" {
		t.Error("CEP should be 12345678")
	}

	if address.State != "SP" {
		t.Error("State should be SP")
	}

	if address.City != "São Paulo" {
		t.Error("City should be São Paulo")
	}

	if address.Neighborhood != "Vila Mariana" {
		t.Error("Neighborhood should be Vila Mariana")
	}

	if address.Street != "Rua Domingos de Morais" {
		t.Error("Street should be Rua Domingos de Morais")
	}

}

func TestBrasilAddressCreationWithEmptyValues(t *testing.T) {

	address := BrasilAPIAddress{}

	if address.CEP != "" {
		t.Error("CEP should be empty")
	}

	if address.State != "" {
		t.Error("State should be empty")
	}

	if address.City != "" {
		t.Error("City should be empty")
	}

	if address.Neighborhood != "" {
		t.Error("Neighborhood should be empty")
	}

	if address.Street != "" {
		t.Error("Street should be empty")
	}

}

func TestBrasilAddressMarshal(t *testing.T) {

	address := BrasilAPIAddress{
		CEP:          "12345678",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Vila Mariana",
		Street:       "Rua Domingos de Morais",
	}

	jsonData, err := json.Marshal(address)

	if err != nil {
		t.Errorf("error marshalling  JSON: %v", err)
	}

	expectedJSON := `{"cep":"12345678","state":"SP","city":"São Paulo","neighborhood":"Vila Mariana","street":"Rua Domingos de Morais"}`
	if string(jsonData) != expectedJSON {
		t.Errorf("expected JSON: %s, got: %s", expectedJSON, string(jsonData))
	}
}

func TestBrasilAddressUnmarshal(t *testing.T) {
	jsonData := []byte(`{"cep":"12345678","state":"SP","city":"São Paulo","neighborhood":"Vila Mariana","street":"Rua Domingos de Morais"}`)

	var address BrasilAPIAddress
	err := json.Unmarshal(jsonData, &address)

	if err != nil {
		t.Errorf("error unmarshalling JSON: %v", err)
	}

	if address.CEP != "12345678" {
		t.Error("CEP should be 12345678")
	}

	if address.State != "SP" {
		t.Error("State should be SP")
	}

	if address.City != "São Paulo" {
		t.Error("City should be São Paulo")
	}

	if address.Neighborhood != "Vila Mariana" {
		t.Error("Neighborhood should be Vila Mariana")
	}

	if address.Street != "Rua Domingos de Morais" {
		t.Error("Street should be Rua Domingos de Morais")
	}
}
