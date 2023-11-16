package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	Cotacao struct {
		Usdbrl Usdbrl `json:"USDBRL"`
	}

	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
)

func main() {
	http.HandleFunc("/cotacao", BuscaCotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()
	cotacao, err := BuscaCotacao(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	usdbrl := convertToEntity(cotacao)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err = insertDatabase(usdbrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usdbrl.Bid)
}

func BuscaCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	return &cotacao, nil
}

func insertDatabase(usdbrl *Usdbrl) error {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&Usdbrl{})
	db.Create(usdbrl)
	return nil
}

func convertToEntity(cotacao *Cotacao) *Usdbrl {
	return &Usdbrl{
		Code:       cotacao.Usdbrl.Code,
		Codein:     cotacao.Usdbrl.Codein,
		Name:       cotacao.Usdbrl.Name,
		High:       cotacao.Usdbrl.High,
		Low:        cotacao.Usdbrl.Low,
		VarBid:     cotacao.Usdbrl.VarBid,
		PctChange:  cotacao.Usdbrl.PctChange,
		Bid:        cotacao.Usdbrl.Bid,
		Ask:        cotacao.Usdbrl.Ask,
		Timestamp:  cotacao.Usdbrl.Timestamp,
		CreateDate: cotacao.Usdbrl.CreateDate,
	}
}
