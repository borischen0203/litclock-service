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

//Should return 400 with "Invalid input" when input hour is over 23.
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

//Should return 400 with "Invalid input" when input hour is less than 00.
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

//Should return 400 with "Invalid input" when input minute is over 59.
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

//Should return 400 with "Invalid input" when input minute is less than 00.
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

//Should return 400 with "Invalid input" when input hour and minute are invalid.
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

//Should return 400 with "Invalid input" when when input second field.
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

//Should return 400 with "Invalid input" when when input uses the operator.
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

//Should return 400 with "Invalid input" when input is not valid time formate.
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

//Should return 400 with "Invalid input" when input are symbols.
func TestInputInvalidType2(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "!@#$%:^&*()",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 400 with "Invalid input" when NumbericTime field is empty string.
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

//Should return 400 with "Invalid input" when missing NumbericTime field.
func TestMissingField(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute is 00.
func TestInputMinuteZero(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:00",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "One o'clock",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute is 01.
func TestInputMinuteOne(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:01",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "One past one",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute 59.
func TestInputMinute59(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:59",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "One to two",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute 15.
func TestInputMinute15(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:15",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Quarter past one",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute 30.
func TestInputMinute30(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:30",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Half past one",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute 45.
func TestInputMinutes45(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:45",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Quarter to two",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute more than 30.
func TestInputMinuteMoreThan30(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:31",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Twenty nine to two",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute less than 29.
func TestInputMinuteLess30(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:29",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Twenty nine past one",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute less than quater.
func TestInputMinuteLess15(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:14",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Fourteen past one",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with human text when input minute more than quater.
func TestInputMinuteLess16(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "13:16",
	}

	responseBody := dto.TimeConvertResponse{
		TextTime: "Sixteen past one",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return the same human text when input hour in 1 and 13.
func TestInputHour13and01(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	requestBody1 := dto.TimeConvertRequest{
		NumbericTime: "01:01",
	}

	requestBody2 := dto.TimeConvertRequest{
		NumbericTime: "13:01",
	}

	request, _ := json.Marshal(requestBody1)
	expected, _ := json.Marshal(requestBody2)

	req1, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))
	req2, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(expected)))

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, w2.Body.String(), w1.Body.String())
}

//Should return the same human text when input hour is 12 and 00.
func TestInputHour12and00(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	requestBody1 := dto.TimeConvertRequest{
		NumbericTime: "12:01",
	}

	requestBody2 := dto.TimeConvertRequest{
		NumbericTime: "00:01",
	}

	request, _ := json.Marshal(requestBody1)
	expected, _ := json.Marshal(requestBody2)

	req1, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))
	req2, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(expected)))

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, w2.Body.String(), w1.Body.String())
}

//Should return the same human text when input hour is 12 and 0.
func TestInputHour12and0(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	requestBody1 := dto.TimeConvertRequest{
		NumbericTime: "12:00",
	}

	requestBody2 := dto.TimeConvertRequest{
		NumbericTime: "0:00",
	}

	request, _ := json.Marshal(requestBody1)
	expected, _ := json.Marshal(requestBody2)

	req1, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))
	req2, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(expected)))

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, w2.Body.String(), w1.Body.String())
}

//Should return 400 with "Invalid input" when input min 0.
func TestInputMin0(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.TimeConvertRequest{
		NumbericTime: "0:0",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidInputError)

	req, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

// //Should return the same human text when input hour is 01 and 1.
func TestInputHour01and1(t *testing.T) {
	logger.Setup()
	config.Setup()

	router := router.SetupRouter()
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	requestBody1 := dto.TimeConvertRequest{
		NumbericTime: "01:00",
	}

	requestBody2 := dto.TimeConvertRequest{
		NumbericTime: "1:00",
	}

	request, _ := json.Marshal(requestBody1)
	expected, _ := json.Marshal(requestBody2)

	req1, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(request)))
	req2, _ := http.NewRequest(http.MethodPost, ConvertTime, strings.NewReader(string(expected)))

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, w2.Body.String(), w1.Body.String())
}
