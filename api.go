package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func NewExchange(key string) *exchange {
	return &exchange{Key: key}
}

type exchange struct {
	Key string
}

type ExchangeRateResponse struct {
	Result         string             `json:"result"`
	Documentation  string             `json:"documentation"`
	TermsOfUse     string             `json:"terms_of_use"`
	Year           int                `json:"year"`
	Month          int                `json:"month"`
	Day            int                `json:"day"`
	BaseCode       string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

// This will return all the exchange rates we have data for on the date 
// in question, in terms of the currency you supplied as the base currency
func (e *exchange) History(baseCurrency string, date time.Time) (ExchangeRateResponse, error) {
	var resData ExchangeRateResponse
	day, month, year := date.Day(), date.Month(), date.Year()
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%v/history/%v/%d/%d/%d", e.Key, baseCurrency, day, month, year)
	res, err := http.Get(url)
	if err != nil {
		return ExchangeRateResponse{}, err
	}
	res.Body.Close()
	reader := bufio.NewReader(res.Body)
	var totalData []byte
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return ExchangeRateResponse{}, err 
		}
		totalData = append(totalData, data...)
	}
	err = json.Unmarshal(totalData, &resData)
	if err != nil {
		return ExchangeRateResponse{}, err
	}
	return resData, nil
}