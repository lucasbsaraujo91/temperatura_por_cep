package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeather(t *testing.T) {
	// Criando uma instância da estrutura Weather
	weather := Weather{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.15,
	}

	// Verificando se os valores estão corretos
	assert.Equal(t, 25.0, weather.TempC, "TempC should be 25.0")
	assert.Equal(t, 77.0, weather.TempF, "TempF should be 77.0")
	assert.Equal(t, 298.15, weather.TempK, "TempK should be 298.15")
}
