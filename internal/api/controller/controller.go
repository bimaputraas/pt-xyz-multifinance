package controller

import "xyz-multifinance/internal/logic"

type (
	Controller struct {
		logic.Logic
	}
)

func New() (*Controller, error) {
	return &Controller{}, nil
}
