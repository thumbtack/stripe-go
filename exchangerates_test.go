package stripe

import (
	"encoding/json"
	"testing"
)

func TestExchangeRatesUnmarshal(t *testing.T) {
	exchangeRatesData := map[string]interface{}{
		"object": "exchange_rates",
		"data": map[string]interface{}{
			"eur": map[string]interface{}{"gbp": 0.891864, "usd": 1.17574},
			"gbp": map[string]interface{}{"eur": 1.12125, "usd": 1.3183},
			"usd": map[string]interface{}{"eur": 0.850525, "gbp": 0.758553},
		},
	}

	bytes, err := json.Marshal(&exchangeRatesData)
	if err != nil {
		t.Error(err)
	}

	var exchangeRates ExchangeRates
	err = json.Unmarshal(bytes, &exchangeRates)
	if err != nil {
		t.Error(err)
	}

	if exchangeRates.Data == nil {
		t.Errorf("Problem deserializing data, got nothing.")
	}

	if exchangeRates.Data["eur"] == nil {
		t.Errorf("Problem deserializing data[eur], got nothing.")
	}

	if exchangeRates.Data["eur"]["gbp"] != 0.891864 {
		t.Errorf("Problem deserializing data[eur][gbp], got %v", exchangeRates.Data["eur"]["gbp"])
	}
}
