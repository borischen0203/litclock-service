package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/borischen0203/litclock-service/config"
	"github.com/borischen0203/litclock-service/dto"
	"github.com/borischen0203/litclock-service/errors"
	"github.com/borischen0203/litclock-service/logger"
	"github.com/borischen0203/litclock-service/router"
	"github.com/stretchr/testify/assert"
)

var ConvertTime = "/api/litclock-service/v1/numeric-time"

func TestHealth(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestVersion(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

//Should return "Invalid input" when input hour is over 23.
func TestInputInvalidTime1(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "24:00",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return "Invalid input" when input hour is less 00.
func TestInputInvalidTime2(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "-01:00",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return "Invalid input" when input minute is over 59.
func TestInputInvalidTime3(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "1:60",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return "Invalid input" when input minute is less than 00.
func TestInputInvalidTime4(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "1:-1",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestInputInvalidTime5(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "24:60",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return "Invalid input" when when input seconds field.
func TestInputInvalidTime6(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "01:01:05",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return "Invalid input" when when input uses the operator.
func TestInputInvalidTime7(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "1+1:1+1",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return "Invalid input" when input is not time type.
func TestInputInvalidType1(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "Hello World",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestInputInvalidType2(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "!@#$%^&*()",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestInputEmptyString(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

// func TestInput(t *testing.T) {
// 	logger.Setup()
// 	config.Setup()

// 	router := router.SetupRouter()
// 	w := httptest.NewRecorder()

// 	requestBody := dto.TimeConvertRequest{
// 		NumbericTime: "0:05",
// 	}

// 	responseBody := dto.TimeConvertResponse{
// 		TextTime: "Five past twelve",
// 	}

// 	request, _ := json.Marshal(requestBody)
// 	expected, _ := json.Marshal(responseBody)

// 	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, string(expected), w.Body.String())
// }
