package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davex98/nobl9-backend/random-generator/application"
	"github.com/davex98/nobl9-backend/random-generator/controllers"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidLengthParam(t *testing.T) {
	service := application.NumberService{}
	server := controllers.NewHttpServer(service)

	testServer := httptest.NewServer(controllers.HandlerFromMux(server, chi.NewMux()))
	requestUrl := fmt.Sprintf("%s/random/mean?requests=1&length=0", testServer.URL)
	get, err := http.Get(requestUrl)

	var errorResponse controllers.Error
	err = json.NewDecoder(get.Body).Decode(&errorResponse)
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, get.StatusCode)
	require.Equal(t, controllers.ErrInvalidLengthParam.Error(), *errorResponse.Error)
}

func TestInvalidRequestsParam(t *testing.T) {
	service := application.NumberService{}
	server := controllers.NewHttpServer(service)

	testServer := httptest.NewServer(controllers.HandlerFromMux(server, chi.NewMux()))
	requestUrl := fmt.Sprintf("%s/random/mean?requests=0&length=1", testServer.URL)
	get, err := http.Get(requestUrl)

	var errorResponse controllers.Error
	err = json.NewDecoder(get.Body).Decode(&errorResponse)
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, get.StatusCode)
	require.Equal(t, controllers.ErrInvalidRequestParam.Error(), *errorResponse.Error)
}

func TestValidRequest(t *testing.T) {
	randomService := randomStub{
		response: []application.RandomNumbersResponse{
			{
				Data: []int{1, 1, 1},
			},
		},
		err: nil,
	}
	service := application.NewNumberService(randomService)
	server := controllers.NewHttpServer(service)
	testServer := httptest.NewServer(controllers.HandlerFromMux(server, chi.NewMux()))

	requestUrl := fmt.Sprintf("%s/random/mean?requests=1&length=3", testServer.URL)
	get, err := http.Get(requestUrl)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, get.StatusCode)

	var response controllers.GeneratedResponse
	err = json.NewDecoder(get.Body).Decode(&response)
	require.NoError(t, err)

	require.Equal(t, 2, len(response))
	require.Equal(t, float32(0), *response[0].Stddev)
	require.Equal(t, 3, len(*response[0].Data))
	require.Equal(t, float32(0), *response[1].Stddev)
	require.Equal(t, 3, len(*response[1].Data))
}

func TestValidMultipleRequest(t *testing.T) {
	randomService := randomStub{
		response: []application.RandomNumbersResponse{
			{Data: []int{1, 1, 1}},
			{Data: []int{1, 2, 3}},
		},
		err: nil,
	}
	service := application.NewNumberService(randomService)
	server := controllers.NewHttpServer(service)
	testServer := httptest.NewServer(controllers.HandlerFromMux(server, chi.NewMux()))

	requestUrl := fmt.Sprintf("%s/random/mean?requests=2&length=3", testServer.URL)
	get, err := http.Get(requestUrl)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, get.StatusCode)

	var response controllers.GeneratedResponse
	err = json.NewDecoder(get.Body).Decode(&response)
	require.NoError(t, err)

	require.Equal(t, 3, len(response))
	require.Equal(t, 3, len(*response[0].Data))
	require.Equal(t, 3, len(*response[1].Data))
	require.Equal(t, 6, len(*response[2].Data))
	require.Equal(t, float32(0), *response[0].Stddev)
	require.Equal(t, float32(0.82), *response[1].Stddev)
	require.Equal(t, float32(0.91), *response[2].Stddev)
}

type randomStub struct {
	response []application.RandomNumbersResponse
	err      error
}

func (r randomStub) GetRandomNumbers(ctx context.Context, length, requestsNumber int) ([]application.RandomNumbersResponse, error) {
	return r.response, r.err
}
