package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CurrencyST struct {
	BaseCode string             `json:"base_code"`
	Rates    map[string]float64 `json:"rates"`
}

func GetCurrency() (*CurrencyST, error) {

	url := "https://exchangerate-api.p.rapidapi.com/rapid/latest/USD"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "19db0dd960mshedd03955ef2c549p136b36jsn88428c70c684")
	req.Header.Add("X-RapidAPI-Host", "exchangerate-api.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	currency := CurrencyST{}
	if ex := json.Unmarshal(body, &currency); ex != nil {
		return nil, ex
	}

	return &currency, nil
}
