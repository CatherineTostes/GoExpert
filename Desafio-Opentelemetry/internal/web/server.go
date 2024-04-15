package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"go.opentelemetry.io/otel"
)

type (
	CepInput struct {
		Cep string `json:"cep"`
	}

	ViaCep struct {
		Cep        string `json:"cep"`
		Localidade string `json:"localidade"`
	}

	Weather struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	Temperature struct {
		City  string  `json:"city"`
		TempC float64 `json:"temp_C"`
		TempF float64 `json:"temp_F"`
		TempK float64 `json:"temp_K"`
	}
)

func GetCep(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(os.Getenv("OTEL_SERVICE_NAME")).Start(r.Context(), "GetCepAndTemp")
	defer span.End()

	var cepInput CepInput
	err := json.NewDecoder(r.Body).Decode(&cepInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if cepInput.Cep == "" || len(cepInput.Cep) == 0 {
		http.Error(w, "zipcode not specified", http.StatusBadRequest)
		return
	}

	if !isValidCep(cepInput.Cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cep, err := GetZipCode(ctx, cepInput.Cep)
	if err != nil || cep.Cep == "" {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := GetWeather(ctx, cep.Localidade)
	if err != nil {
		http.Error(w, "can not find weather", http.StatusNotFound)
		return
	}

	temperature := Temperature{
		City:  cep.Localidade,
		TempC: weather.Current.TempC,
		TempF: weather.Current.TempC*1.8 + 32,
		TempK: weather.Current.TempC + 273.15,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}

func GetZipCode(ctx context.Context, cepParam string) (*ViaCep, error) {
	ctx, span := otel.Tracer(os.Getenv("OTEL_SERVICE_NAME")).Start(ctx, "GetCep")
	defer span.End()
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

func GetWeather(ctx context.Context, localidade string) (*Weather, error) {
	_, span := otel.Tracer(os.Getenv("OTEL_SERVICE_NAME")).Start(ctx, "GetWeather")
	defer span.End()

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
