package service

import "github.com/gorkagg10/equity-calculator-api/internal/domain"

type Asset struct{}

func NewAsset() *Asset {
	return &Asset{}
}

func (a *Asset) Add() (*domain.Asset, error) {
	return nil, nil
}
