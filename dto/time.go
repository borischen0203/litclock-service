package dto

import (
	"time"

	e "github.com/borischen0203/litclock-service/errors"
)

/**
 * This struct creates a time convert request from
 * the one paramater: numberic time
 *
 * @param NumbericTime      clocks time that display as numbers
 */
type TimeConvertRequest struct {
	NumbericTime string `json:"numbericTime" binding:"required"`
}

/**
 * This struct creates a time convert response from
 * the one paramater: text time
 *
 * @param TextTime      clocks time that display as text
 */
type TimeConvertResponse struct {
	TextTime string `json:"textTime" binding:"required"`
}

//Validate function validate the request
func (r TimeConvertRequest) Validate() (bool, e.ErrorInfo) {
	if r.IsEmpty() {
		return false, e.InvalidInputError
	}
	if !r.IsValidTimeFormate() {
		return false, e.InvalidInputError
	}
	return true, e.NoError
}

//IsValidAlias function validate Alias is valid alphabet or number, and less than 30 lengths
func (r TimeConvertRequest) IsValidTimeFormate() bool {
	_, err := time.Parse("15:04", r.NumbericTime)
	return err == nil
}

//IsEmpty function validate long url is not empty input
func (r TimeConvertRequest) IsEmpty() bool {
	return r.NumbericTime == ""
}
