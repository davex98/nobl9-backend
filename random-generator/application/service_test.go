package application_test

import (
	"context"
	"github.com/davex98/nobl9-backend/random-generator/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeviation(t *testing.T) {
	tests := []struct {
		input []int
		want  float32
	}{
		{input: []int{1, 2, 3, 4, 5}, want: 1.41},
		{input: []int{}, want: 0},
		{input: []int{1, -1}, want: 1},
		{input: []int{1, 1, 1, 1, 1}, want: 0},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("Get deviation test", func(t *testing.T) {
			t.Parallel()
			deviation := application.GetDeviation(tc.input...)

			if deviation != tc.want {
				t.Errorf("expected %v, got %v", tc.want, deviation)
			}
		})
	}
}

func TestServerFlow(t *testing.T) {
	ctx := context.Background()
	randomService := randomStub{
		response: []application.RandomNumbersResponse{
			{
				Data: []int{1, 1, 1},
			},
		},
		err: nil,
	}
	app := application.NewNumberService(randomService)

	numbers, err := app.CollectRandomNumbers(ctx, 3, 1)
	require.NoError(t, err)

	deviation := float32(0)
	require.Equal(t, len(numbers), 2)
	require.Equal(t, deviation, numbers[0].Stddev)
	require.Equal(t, deviation, numbers[1].Stddev)
}

func TestServerBiggerFlow(t *testing.T) {
	ctx := context.Background()
	randomService := randomStub{
		response: []application.RandomNumbersResponse{
			{Data: []int{1, 1, 1}},
			{Data: []int{1, 2, 3, 4, 5}},
		},
		err: nil,
	}
	app := application.NewNumberService(randomService)

	numbers, err := app.CollectRandomNumbers(ctx, 3, 1)
	require.NoError(t, err)

	require.Equal(t, len(numbers), 3)
	require.Equal(t, float32(0), numbers[0].Stddev)
	require.Equal(t, float32(1.41), numbers[1].Stddev)
	require.Equal(t, float32(1.5), numbers[2].Stddev)
}

type randomStub struct {
	response []application.RandomNumbersResponse
	err      error
}

func (r randomStub) GetRandomNumbers(ctx context.Context, length, requestsNumber int) ([]application.RandomNumbersResponse, error) {
	return r.response, r.err
}
