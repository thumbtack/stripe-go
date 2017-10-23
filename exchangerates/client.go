// Package exchangerates provides the /exchange_rates APIs
package exchangerates

import (
	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /exchange_rates and exchangerates-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns an ExchangeRates for a given currency
func Get(currency string) (*stripe.ExchangeRates, error) {
	return getC().Get(currency)
}

func (c Client) Get(currency string) (*stripe.ExchangeRates, error) {
	exchangeRates := &stripe.ExchangeRates{}
	err := c.B.Call("GET", "/exchange_rates/"+currency, c.Key, nil, nil, exchangeRates)

	return exchangeRates, err
}

// List returns an ExchangeRates for all currencies
func List() (*stripe.ExchangeRates, error) {
	return getC().List()
}

func (c Client) List() (*stripe.ExchangeRates, error) {
	exchangeRates := &stripe.ExchangeRates{}
	err := c.B.Call("GET", "/exchange_rates", c.Key, nil, nil, exchangeRates)

	return exchangeRates, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
