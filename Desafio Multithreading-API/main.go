package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	cep := "01153000"
	brasilAPI := make(chan string)
	viaCep := make(chan string)

	go func() {
		firstResponse, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(firstResponse.Body)
		if err != nil {
			panic(err)
		}
		brasilAPI <- string(body)
		defer firstResponse.Body.Close()
	}()
	go func() {
		secondResponse, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(secondResponse.Body)
		if err != nil {
			panic(err)
		}
		viaCep <- string(body)
		defer secondResponse.Body.Close()
	}()

	select {
	case response := <-brasilAPI:
		println("API: BrasilAPI: " + response)
	case response := <-viaCep:
		println("API: ViaCep" + response)
	case <-time.After(time.Second):
		println("timeout")
	}
}
