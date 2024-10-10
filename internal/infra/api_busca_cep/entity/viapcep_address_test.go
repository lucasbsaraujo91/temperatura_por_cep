package entity

import (
	"encoding/json"
	"testing"
)

func TestViaCepCreation(t *testing.T) {

	address := ViaCEPAddress{
		CEP:         "12345678",
		Logradouro:  "Rua Domingos de Morais",
		Complemento: "Complemento",
		Bairro:      "Vila Mariana",
		Localidade:  "São Paulo",
		UF:          "SP",
		IBGE:        "123456",
		GIA:         "1234",
		DDD:         "11",
		SIAFI:       "1234",
	}

	if address.CEP != "12345678" {
		t.Error("CEP should be 12345678")
	}

	if address.Logradouro != "Rua Domingos de Morais" {
		t.Error("Logradouro should be Rua Domingos de Morais")
	}

	if address.Complemento != "Complemento" {
		t.Error("Complemento should be Complemento")
	}

	if address.Bairro != "Vila Mariana" {
		t.Error("Bairro should be Vila Mariana")
	}

	if address.Localidade != "São Paulo" {
		t.Error("Localidade should be São Paulo")
	}

	if address.UF != "SP" {
		t.Error("UF should be SP")
	}

	if address.IBGE != "123456" {
		t.Error("IBGE should be 123456")
	}

	if address.GIA != "1234" {
		t.Error("GIA should be 1234")
	}

	if address.DDD != "11" {
		t.Error("DDD should be 11")
	}

	if address.SIAFI != "1234" {
		t.Error("SIAFI should be 1234")
	}

}

func TestViaCepCreationWithEmptyValues(t *testing.T) {

	address := ViaCEPAddress{}

	if address.CEP != "" {
		t.Error("CEP should be empty")
	}

	if address.Logradouro != "" {
		t.Error("Logradouro should be empty")
	}

	if address.Complemento != "" {
		t.Error("Complemento should be empty")
	}

	if address.Bairro != "" {
		t.Error("Bairro should be empty")
	}

	if address.Localidade != "" {
		t.Error("Localidade should be empty")
	}

	if address.UF != "" {
		t.Error("UF should be empty")
	}

	if address.IBGE != "" {
		t.Error("IBGE should be empty")
	}

	if address.GIA != "" {
		t.Error("GIA should be empty")
	}

	if address.DDD != "" {
		t.Error("DDD should be empty")
	}

	if address.SIAFI != "" {
		t.Error("SIAFI should be empty")
	}

}

func TestViaCepAddressMarshal(t *testing.T) {

	address := ViaCEPAddress{
		CEP:         "12345678",
		Logradouro:  "Rua Domingos de Morais",
		Complemento: "Complemento",
		Bairro:      "Vila Mariana",
		Localidade:  "São Paulo",
		UF:          "SP",
		IBGE:        "123456",
		GIA:         "1234",
		DDD:         "11",
		SIAFI:       "1234",
	}

	jsonData, err := json.Marshal(address)

	if err != nil {
		t.Errorf("error marshalling  JSON: %v", err)
	}

	expectedJSON := `{"cep":"12345678","logradouro":"Rua Domingos de Morais","complemento":"Complemento","bairro":"Vila Mariana","localidade":"São Paulo","uf":"SP","ibge":"123456","gia":"1234","ddd":"11","siafi":"1234"}`
	if string(jsonData) != expectedJSON {
		t.Errorf("JSON was incorrect, got: %s, want: %s.", jsonData, expectedJSON)
	}
}

func TestViaCepAddressUnmarshal(t *testing.T) {

	jsonData := []byte(`{"cep":"12345678","logradouro":"Rua Domingos de Morais","complemento":"Complemento","bairro":"Vila Mariana","localidade":"São Paulo","uf":"SP","ibge":"123456","gia":"1234","ddd":"11","siafi":"1234"}`)

	var address ViaCEPAddress
	err := json.Unmarshal(jsonData, &address)

	if err != nil {
		t.Errorf("error unmarshalling JSON: %v", err)
	}

	if address.CEP != "12345678" {
		t.Error("CEP should be 12345678")
	}

	if address.Logradouro != "Rua Domingos de Morais" {
		t.Error("Logradouro should be Rua Domingos de Morais")
	}

	if address.Complemento != "Complemento" {
		t.Error("Complemento should be Complemento")
	}

	if address.Bairro != "Vila Mariana" {
		t.Error("Bairro should be Vila Mariana")
	}

	if address.Localidade != "São Paulo" {
		t.Error("Localidade should be São Paulo")
	}

	if address.UF != "SP" {
		t.Error("UF should be SP")
	}

	if address.IBGE != "123456" {
		t.Error("IBGE should be 123456")
	}

	if address.GIA != "1234" {
		t.Error("GIA should be 1234")
	}

	if address.DDD != "11" {
		t.Error("DDD should be 11")
	}

	if address.SIAFI != "1234" {
		t.Error("SIAFI should be 1234")
	}

}
