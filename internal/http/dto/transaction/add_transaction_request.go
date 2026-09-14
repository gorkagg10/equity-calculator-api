package transaction

type AddTransactionRequest struct {
	Symbol string `json:"symbol"`
	Shares int    `json:"shares"`
}
