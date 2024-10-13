package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherResponseInitialization(t *testing.T) {
	// Testando a inicialização da estrutura WeatherResponse
	weather := WeatherResponse{
		Location: Location{
			Name:        "São Paulo",
			Region:      "SP",
			Country:     "Brasil",
			Lat:         -23.5505,
			Lon:         -46.6333,
			TzID:        "America/Sao_Paulo",
			Localtime:   "2024-10-13 14:00",
			LocaltimeEp: 1697205600,
		},
		Current: Current{
			LastUpdatedEp: 1697205600,
			LastUpdated:   "2024-10-13 14:00",
			TempC:         25.5,
			TempF:         77.9,
			TempK:         298.65,
			IsDay:         1,
			Condition: Condition{
				Text: "Sunny",
				Icon: "//cdn.weatherapi.com/weather/64x64/day/113.png",
				Code: 1000,
			},
			WindMph:      10.5,
			WindKph:      16.9,
			WindDegree:   180,
			WindDir:      "South",
			PressureMb:   1013,
			PressureIn:   29.91,
			PrecipMm:     0.0,
			PrecipIn:     0.0,
			Humidity:     60,
			Cloud:        0,
			FeelsLikeC:   26.0,
			FeelsLikeF:   78.8,
			WindChillC:   24.0,
			WindChillF:   75.2,
			HeatIndexC:   26.5,
			HeatIndexF:   79.7,
			DewPointC:    15.0,
			DewPointF:    59.0,
			VisibilityKm: 10.0,
			VisibilityMi: 6.2,
			UV:           5.0,
			GustMph:      15.0,
			GustKph:      24.1,
		},
	}

	assert.Equal(t, "São Paulo", weather.Location.Name)
	assert.Equal(t, "SP", weather.Location.Region)
	assert.Equal(t, "Brasil", weather.Location.Country)
	assert.Equal(t, 25.5, weather.Current.TempC)
	assert.Equal(t, "Sunny", weather.Current.Condition.Text)
}

func TestLocationInitialization(t *testing.T) {
	// Testando a inicialização da estrutura Location
	location := Location{
		Name:        "Rio de Janeiro",
		Region:      "RJ",
		Country:     "Brasil",
		Lat:         -22.9068,
		Lon:         -43.1729,
		TzID:        "America/Sao_Paulo",
		Localtime:   "2024-10-13 14:00",
		LocaltimeEp: 1697205600,
	}

	assert.Equal(t, "Rio de Janeiro", location.Name)
	assert.Equal(t, "RJ", location.Region)
	assert.Equal(t, "Brasil", location.Country)
	assert.Equal(t, -22.9068, location.Lat)
	assert.Equal(t, -43.1729, location.Lon)
}

func TestCurrentInitialization(t *testing.T) {
	// Testando a inicialização da estrutura Current
	current := Current{
		LastUpdatedEp: 1697205600,
		LastUpdated:   "2024-10-13 14:00",
		TempC:         25.5,
		IsDay:         1,
		Condition: Condition{
			Text: "Clear",
			Icon: "//cdn.weatherapi.com/weather/64x64/day/113.png",
			Code: 1000,
		},
	}

	assert.Equal(t, "Clear", current.Condition.Text)
	assert.Equal(t, 25.5, current.TempC)
	assert.Equal(t, 1, current.IsDay)
}

func TestConditionInitialization(t *testing.T) {
	// Testando a inicialização da estrutura Condition
	condition := Condition{
		Text: "Partly cloudy",
		Icon: "//cdn.weatherapi.com/weather/64x64/day/116.png",
		Code: 1003,
	}

	assert.Equal(t, "Partly cloudy", condition.Text)
	assert.Equal(t, "//cdn.weatherapi.com/weather/64x64/day/116.png", condition.Icon)
	assert.Equal(t, 1003, condition.Code)
}
