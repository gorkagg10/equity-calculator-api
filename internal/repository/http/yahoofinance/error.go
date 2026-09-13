package yahoofinance

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
