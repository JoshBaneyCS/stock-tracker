package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ConvertCurrency(amount float64, from, to string) (float64, error) {
	url := fmt.Sprintf("https://api.exchangerate.host/convert?from=%s&to=%s&amount=%.2f", from, to, amount)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Result float64 `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Result, nil
}
