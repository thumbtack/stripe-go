package stripe

// ExchangeRates is the resource representing the currency exchange rates at
// a given time.
type ExchangeRates struct {
	Data map[Currency]map[Currency]float64 `json:"data"`
}
