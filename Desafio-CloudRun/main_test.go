package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCep(t *testing.T) {
	t.Run("should return true when cep is valid", func(t *testing.T) {
		cep := "22250040"
		valid := isValidCep(cep)
		assert.True(t, valid)
	})

	t.Run("should return false when cep is invalid", func(t *testing.T) {
		cepNumberInvalid := "2225aa40"
		cepLengthInvalid := "2225004"

		assert.False(t, isValidCep(cepNumberInvalid))

		assert.False(t, isValidCep(cepLengthInvalid))
	})
}

func TestGetZipCode(t *testing.T) {
	cep := "22250040"
	address, err := GetZipCode(cep)
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, "RJ", address.Uf)
}

func TestGetWeather(t *testing.T) {
	weather, err := GetWeather("Rio de Janeiro")
	assert.Nil(t, err)
	assert.NotNil(t, weather)
	assert.NotEmpty(t, weather.Current.TempC)
}
