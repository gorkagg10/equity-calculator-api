package domain

type AssetData struct {
	symbol   string
	name     string
	price    float64
	currency string
	exchange string
}

func NewAssetData(
	symbol string,
	name string,
	price float64,
	currency string,
	exchange string,
) *AssetData {
	return &AssetData{
		symbol:   symbol,
		name:     name,
		price:    price,
		currency: currency,
		exchange: exchange,
	}
}

func (d AssetData) Symbol() string {
	return d.symbol
}

func (d AssetData) Name() string {
	return d.name
}

func (d AssetData) Price() float64 {
	return d.price
}

func (d AssetData) Currency() string {
	return d.currency
}

func (d AssetData) Exchange() string {
	return d.exchange
}
