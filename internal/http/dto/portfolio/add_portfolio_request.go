package portfolio

import "errors"

type AddPortfolioRequest struct {
	Name string `json:"name"`
}

func (a *AddPortfolioRequest) Validate() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
