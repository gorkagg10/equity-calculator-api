package domain

import "errors"

var (
	ErrPortfolioNotFound = errors.New("portfolio not found")
	ErrAssetNotFound     = errors.New("asset not found")
	ErrPositionNotFound  = errors.New("position not found")
)
