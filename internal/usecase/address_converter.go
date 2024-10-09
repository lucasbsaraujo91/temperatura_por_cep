package usecase

import "temperatura_por_cep/internal/infra/api/entity"

type Address struct {
	ZipCode      string
	Street       string
	Neighborhood string
	City         string
	State        string
	Complemento  string
	IBGE         string
	GIA          string
	DDD          string
	SIAFI        string
}

func ConvertBrasilAPIAddress(brazilAddress entity.BrasilAPIAddress) *Address {
	return &Address{
		ZipCode:      brazilAddress.CEP,
		Street:       brazilAddress.Street,
		Neighborhood: brazilAddress.Neighborhood,
		City:         brazilAddress.City,
		State:        brazilAddress.State,
	}
}

func ConvertViaCEPAddress(viaCEPAddress entity.ViaCEPAddress) *Address {
	return &Address{
		ZipCode:      viaCEPAddress.CEP,
		Street:       viaCEPAddress.Logradouro,
		Neighborhood: viaCEPAddress.Bairro,
		City:         viaCEPAddress.Localidade,
		State:        viaCEPAddress.UF,
		Complemento:  viaCEPAddress.Complemento,
		IBGE:         viaCEPAddress.IBGE,
		GIA:          viaCEPAddress.GIA,
		DDD:          viaCEPAddress.DDD,
		SIAFI:        viaCEPAddress.SIAFI,
	}
}
