package asset

import "errors"

type AddAssetRequest struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Category string `json:"category"`
}

func (c *AddAssetRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Symbol == "" {
		return errors.New("symbol is required")
	}
	if c.Category == "" {
		return errors.New("category is required")
	}
	return nil
}
