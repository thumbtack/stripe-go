package stripe

// SourceTransactionListParams is the set of parameters that can be used when listing SourceTransactions.
type SourceTransactionListParams struct {
	ListParams `form:"*"`
	Source     string `form:"-"` // Sent in with the URL
}

// SourceTransactionList is a list object for SourceTransactions.
type SourceTransactionList struct {
	ListMeta
	Values []*SourceTransaction `json:"data"`
}

// SourceTransaction is the resource representing a Stripe source transaction.
type SourceTransaction struct {
	Amount       int64    `json:"amount"`
	Created      int64    `json:"created"`
	Currency     Currency `json:"currency"`
	CustomerData string   `json:"customer_data"`
	ID           string   `json:"id"`
	Live         bool     `json:"livemode"`
	Source       string   `json:"source"`
	Type         string   `json:"type"`
}
