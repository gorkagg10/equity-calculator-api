package yahoofinance

type YahooChartResponse struct {
	Chart struct {
		Result []struct {
			Meta AssetData `json:"meta"`
		} `json:"result"`
		Error *Error `json:"error"`
	} `json:"chart"`
}
