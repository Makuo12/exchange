package api


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

type ExchangeRateCurrentResponse struct {
	Result              string             `json:"result"`
	Documentation       string             `json:"documentation"`
	TermsOfUse          string             `json:"terms_of_use"`
	TimeLastUpdateUnix  int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC   string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix  int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC   string             `json:"time_next_update_utc"`
	BaseCode            string             `json:"base_code"`
	ConversionRates     map[string]float64 `json:"conversion_rates"`
}