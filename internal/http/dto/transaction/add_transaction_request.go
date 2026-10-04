package transaction

type AddTransactionRequest struct {
	AssetID         string  `json:"asset_id"`
	TransactionType string  `json:"transaction_type"`
	Quantity        float64 `json:"quantity"`
	UnitPrice       float64 `json:"unit_price"`
	Currency        string  `json:"currency"`
}
