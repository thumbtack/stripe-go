package exchangerates

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	_ "github.com/stripe/stripe-go/testing"
)

func TestExchangeRatesGet(t *testing.T) {
	// TODO: uncomment this line once stripe-mock supports /v1/exchange_rates
	t.Skip("not yet supported by stripe-mock")

	rates, err := Get("usd")
	assert.Nil(t, err)
	assert.NotNil(t, rates)
}

func TestExchangeRatesList(t *testing.T) {
	// TODO: uncomment this line once stripe-mock supports /v1/exchange_rates
	t.Skip("not yet supported by stripe-mock")

	rates, err := List()
	assert.Nil(t, err)
	assert.NotNil(t, rates)
}
