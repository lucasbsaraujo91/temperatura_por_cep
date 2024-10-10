// weather_usecase.go
package usecase

import (
	"fmt"
	"io"
	"strings"
	"temperatura_por_cep/internal/entity"
	"temperatura_por_cep/internal/infra/api_busca_temperatura/service"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type WeatherUseCase struct {
	AddressUseCase *AddressUseCase
	WeatherService *service.WeatherService
}

func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (*entity.Weather, error) {
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

func (u *WeatherUseCase) getWeatherByCity(city string, state string) (*entity.Weather, error) {

	newCity := sanitizeCity(removeAccents(city))

	fullWeather, err := u.WeatherService.FetchWeatherByCity(newCity, state)
	if err != nil {
		return nil, err
	}

	// Converte a resposta para o formato desejado
	weather := &entity.Weather{
		TempC: fullWeather.Current.TempC,
		TempF: (fullWeather.Current.TempC * 1.8) + 32,
		TempK: fullWeather.Current.TempC + 273.15,
	}

	return weather, nil
}

func sanitizeCity(city string) string {
	return strings.ReplaceAll(city, " ", "+")
}

// removeAccents remove acentos de uma string
func removeAccents(str string) string {
	t := transform.NewReader(strings.NewReader(str), norm.NFD)
	normalized, _ := io.ReadAll(t)
	result := string(normalized)

	var sb strings.Builder
	for _, r := range result {
		if unicode.Is(unicode.Mn, r) {
			continue // Ignora caracteres de acento
		}
		sb.WriteRune(r)
	}
	return sb.String()
}
