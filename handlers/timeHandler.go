package handlers

import (
	// "fmt"

	"net/http"

	"github.com/borischen0203/litclock-service/dto"
	"github.com/borischen0203/litclock-service/errors"
	"github.com/borischen0203/litclock-service/logger"
	"github.com/borischen0203/litclock-service/services"

	"github.com/gin-gonic/gin"
)

func ConvertTime(c *gin.Context) {
	request := dto.TimeConvertRequest{}
	response := dto.TimeConvertResponse{}
	c.BindJSON(&request)
	logger.Info.Printf("[ConvertTimeHandler] request=%+v\n", request)
	statusCode, result, err := services.ConvertTimeService(request, response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}
