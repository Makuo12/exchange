package api

import (
	"fmt"
	"testing"
	"time"

)


func TestHistory(t *testing.T) {
	customTime := time.Date(
    2010,           // Year
    time.January,   // Month
    1,              // Day
    0,              // Hour
    0,              // Minute
    0,              // Second
    0,              // Nanosecond
    time.UTC,       // Time zone
	)
	ex := NewExchange("5c8f5a4753ff83ceb1f4f28a")
	data, err := ex.History("USD", customTime)
	if err != nil {
		t.Log(err)
		return
	}
	if data.Result == "error" {
		t.Error("Data is meant to be a success")
	}
	fmt.Println(data)
}

func TestCurrent(t *testing.T) {
	ex := NewExchange("5c8f5a4753ff83ceb1f4f28a")
	data, err := ex.Current("USD")
	if err != nil {
		t.Log(err)
		return
	}
	if data.Result == "error" {
		t.Error("Data is meant to be a success")
	}
	fmt.Println(data)
}