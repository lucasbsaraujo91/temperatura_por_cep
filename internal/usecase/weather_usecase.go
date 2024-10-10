// weather_usecase.go
package usecase

import (
	"fmt"
	"temperatura_por_cep/internal/entity"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/service"
	"temperatura_por_cep/internal/utils"
)

// DTOs de entrada e saída

type GetWeatherOutputDTO struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type WeatherUseCase struct {
	AddressUseCase *AddressUseCase
	WeatherService *service.WeatherService
}

func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (*GetWeatherOutputDTO, error) {
	// Busca o endereço pelo CEP
	address, err := u.getAddressByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	// Busca o clima pela cidade do endereço
	weather, err := u.getWeatherByCity(address.City, address.State)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (u *WeatherUseCase) getAddressByZipCode(zipCode string) (*entity.Address, error) {
	// Cria o DTO de entrada
	input := GetAddressInputDTO{ZipCode: zipCode}

	// Chama o método GetAddressByZipCode
	addressOutput, err := u.AddressUseCase.GetAddressByZipCode(input)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter endereço: %v", err)
	}

	address := &entity.Address{
		Street:  addressOutput.Street,
		City:    addressOutput.City,
		State:   addressOutput.State,
		ZipCode: addressOutput.ZipCode,
	}
	return address, nil
}

func (u *WeatherUseCase) getWeatherByCity(city string, state string) (*GetWeatherOutputDTO, error) {
	newCity := utils.SanitizeCity(utils.RemoveAccents(city))

	fullWeather, err := u.WeatherService.FetchWeatherByCity(newCity, state)
	if err != nil {
		return nil, err
	}

	// Converte a resposta para o formato desejado
	weather := &GetWeatherOutputDTO{
		TempC: fullWeather.Current.TempC,
		TempF: utils.ConvertCelsiusToFahrenheit(fullWeather.Current.TempC),
		TempK: utils.ConvertCelsiusToKelvin(fullWeather.Current.TempC),
	}

	return weather, nil
}
