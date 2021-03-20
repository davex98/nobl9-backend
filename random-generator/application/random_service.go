package application

import (
	"context"
)

type RandomService interface {
	GetRandomNumbers(ctx context.Context, length, requestsNumber int) ([]RandomNumbersResponse, error)
}

type RandomNumbersResponse struct {
	Data []int
}
