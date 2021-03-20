package application

import (
	"context"
	"math"
)

type NumberService struct {
	randomService RandomService
}

func NewNumberService(service RandomService) NumberService {
	return NumberService{randomService: service}
}

func (s NumberService) CollectRandomNumbers(ctx context.Context, length, requestsNumber int) ([]DeviationResponse, error) {
	numbers, err := s.randomService.GetRandomNumbers(ctx, length, requestsNumber)
	if err != nil {
		return nil, err
	}
	var response []DeviationResponse
	var totalData []int
	for _, number := range numbers {
		deviation := GetDeviation(number.Data...)
		totalData = append(totalData, number.Data...)
		response = append(response, DeviationResponse{
			Stddev: deviation,
			Data:   number.Data,
		})
	}

	deviation := GetDeviation(totalData...)
	response = append(response, DeviationResponse{
		Stddev: deviation,
		Data:   totalData,
	})

	return response, nil
}

func GetDeviation(numbers ...int) float32 {
	if len(numbers) == 0 {
		return 0
	}
	total := 0
	for _, number := range numbers {
		total += number
	}
	size := len(numbers)
	mean := total / size

	var sd float64
	for _, number := range numbers {
		a := (number) - mean
		sd += math.Pow(float64(a), 2)
	}
	sd = math.Sqrt(sd / float64(size))
	output := math.Round(sd*100) / 100
	return float32(output)
}

type DeviationResponse struct {
	Stddev float32
	Data   []int
}
