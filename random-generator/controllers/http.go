package controllers

import (
	"errors"
	"github.com/davex98/nobl9-backend/random-generator/application"
	"github.com/go-chi/render"
	"net/http"
)

type server struct {
	app application.NumberService
}

func NewHttpServer(app application.NumberService) ServerInterface {
	return server{app: app}
}

func (s server) GetRandomMean(w http.ResponseWriter, r *http.Request, params GetRandomMeanParams) {
	ok, err := validateParams(params)
	if !ok {
		handleClientError(w, r, err)
		return
	}
	numbers, err := s.app.CollectRandomNumbers(r.Context(), params.Length, params.Requests)
	if err != nil {
		handleServerError(w, r, err)
		return
	}
	response := getResponse(numbers)
	render.Respond(w, r, response)
}

func validateParams(params GetRandomMeanParams) (bool, error) {
	if params.Length <= 0 {
		return false, ErrInvalidLengthParam
	}
	if params.Requests <= 0 {
		return false, ErrInvalidRequestParam
	}
	return true, nil
}

func handleClientError(w http.ResponseWriter, r *http.Request, err error) {
	errorMessage := err.Error()
	resp := Error{
		Error: &errorMessage,
	}
	w.WriteHeader(http.StatusBadRequest)
	render.Respond(w, r, resp)
}

func handleServerError(w http.ResponseWriter, r *http.Request, err error) {
	errorMessage := err.Error()
	resp := Error{
		Error: &errorMessage,
	}
	w.WriteHeader(http.StatusInternalServerError)
	render.Respond(w, r, resp)
}

func getResponse(numbers []application.DeviationResponse) GeneratedResponse {
	var response []Response
	for _, number := range numbers {
		a := number.Data
		b := number.Stddev
		response = append(response, Response{
			Data:   &a,
			Stddev: &b,
		})
	}
	return response
}

var ErrInvalidLengthParam = errors.New("length param has to be bigger than 0")
var ErrInvalidRequestParam = errors.New("request param has to be bigger than 0")
