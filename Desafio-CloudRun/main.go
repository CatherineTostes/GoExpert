package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type (
	ViaCep struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}

	Weather struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	Temperature struct {
		TempC float64 `json:"temp_C"`
		TempF float64 `json:"temp_F"`
		TempK float64 `json:"temp_K"`
	}
)

func main() {
	http.HandleFunc("/", BuscaCepAndWeatherHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepAndWeatherHandler(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" || len(cepParam) == 0 {
		http.Error(w, "zipcode not specified", http.StatusBadRequest)
		return
	}

	if !isValidCep(cepParam) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cep, err := GetZipCode(cepParam)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := GetWeather(cep.Localidade)
	if err != nil {
		http.Error(w, "can not find weather", http.StatusNotFound)
		return
	}

	temperature := Temperature{
		TempC: weather.Current.TempC,
		TempF: weather.Current.TempC*1.8 + 32,
		TempK: weather.Current.TempC + 273.15,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}

func GetWeather(localidade string) (*Weather, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&key=5b2091df03de46afb5014946240904", url.QueryEscape(localidade)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var w Weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func GetZipCode(cepParam string) (*ViaCep, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cepParam))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var v ViaCep
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func isValidCep(cep string) bool {
	if len(cep) != 8 {
		return false
	}

	_, err := strconv.Atoi(cep)
	if err != nil {
		return false
	}

	return true
}
