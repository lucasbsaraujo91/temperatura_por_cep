// weather_usecase.go
package usecase

import (
	"errors"
	"temperatura_por_cep/internal/entity"
)

type WeatherUseCase struct {
	// Adicione os repositórios ou clientes necessários
}

func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (*entity.Weather, error) {
	if len(zipCode) != 8 {
		return nil, errors.New("invalid zipcode")
	}

	// Busca o endereço pelo CEP
	address, err := u.getAddressByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	// Busca o clima pela cidade do endereço
	weather, err := u.getWeatherByCity(address.City)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (u *WeatherUseCase) getAddressByZipCode(zipCode string) (*entity.Address, error) {
	// Implementar chamada à API do ViaCEP
	address := &entity.Address{
		ZipCode:      zipCode,
		Street:       "Rua Exemplo",
		Neighborhood: "Bairro Exemplo",
		City:         "São Paulo",
		State:        "SP",
	}
	return address, nil
}

func (u *WeatherUseCase) getWeatherByCity(city string) (*entity.Weather, error) {
	// Implementar chamada à API WeatherAPI
	weather := &entity.Weather{
		TempC: 28.5,
		TempF: 28.5*1.8 + 32,
		TempK: 28.5 + 273,
	}
	return weather, nil
}
