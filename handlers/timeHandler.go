package handlers

/***
 *
 * This file mainly handle time
 *
 * @author: Boris
 * @version: 2021-10-23
 *
 */

import (
	"net/http"

	"github.com/borischen0203/litclock-service/dto"
	"github.com/borischen0203/litclock-service/errors"
	"github.com/borischen0203/litclock-service/logger"
	"github.com/borischen0203/litclock-service/services"

	"github.com/gin-gonic/gin"
)

// @Summary convert time from humeric to text
// @Description convert numeric time to human friendly text
// @Tags convert time
// @Accept json
// @Produce json
// @Param body body dto.TimeConvertRequest true "body"
// @Success 200 {object} dto.TimeConvertResponse "ok"
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /api/litclock-service/v1/numeric-time [post]
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
