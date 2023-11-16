package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	arq, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer arq.Close()
	_, err = arq.Write([]byte("Dólar: " + string(body)))
	if err != nil {
		log.Println(err)
		panic(err)
	}

}
