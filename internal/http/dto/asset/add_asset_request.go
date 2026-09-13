package asset

import "errors"

type AddAssetRequest struct {
	Symbol string `json:"symbol"`
}

func (c *AddAssetRequest) Validate() error {
	if c.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}
