package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// This will return all the exchange rates we have data for on the date
// in question, in terms of the currency you supplied as the base currency
func (e *exchange) History(baseCurrency string, date time.Time) (ExchangeRateResponse, error) {
	var resData ExchangeRateResponse
	day, month, year := date.Day(), date.Month(), date.Year()
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/history/%s/%d/%d/%d", e.Key, baseCurrency, year, int(month), day)
	res, err := http.Get(url)
	if err != nil {
		return ExchangeRateResponse{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExchangeRateResponse{}, err 
	}
	err = json.Unmarshal(data, &resData)
	if err != nil {
		return ExchangeRateResponse{}, err
	}
	if resData.Result == "error" {
		return ExchangeRateResponse{}, errors.New("Your payment plan needs to be upgraded")
	}
	return resData, nil
}



// This will return all the exchange rates we have 
// data for on the provided base currency
func (e *exchange) Current(baseCurrency string) (ExchangeRateCurrentResponse, error) {
	var resData ExchangeRateCurrentResponse
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", e.Key, baseCurrency)
	res, err := http.Get(url)
	if err != nil {
		return ExchangeRateCurrentResponse{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExchangeRateCurrentResponse{}, err 
	}
	err = json.Unmarshal(data, &resData)
	if err != nil {
		return ExchangeRateCurrentResponse{}, err
	}
	if resData.Result == "error" {
		return ExchangeRateCurrentResponse{}, errors.New("Your payment plan needs to be upgraded")
	}
	return resData, nil
}