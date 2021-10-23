package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/borischen0203/litclock-service/dto"
	e "github.com/borischen0203/litclock-service/errors"
	"github.com/borischen0203/litclock-service/logger"
)

func ConvertTimeService(request dto.TimeConvertRequest, response dto.TimeConvertResponse) (int64, dto.TimeConvertResponse, e.ErrorInfo) {
	if ok, err := request.Validate(); !ok {
		return 400, response, err
	}

	t, err := time.Parse("15:04", request.NumericTime)

	if err != nil {
		logger.Error.Printf("[ConvertTimeService] time parsing error: %v\n", err)
		return 500, response, e.InternalServerError
	}
	humanText := TimeToWords(t.Hour(), t.Minute())
	res := dto.TimeConvertResponse{
		TextTime: humanText,
	}
	return 200, res, e.NoError
}

//This function mainly convert numeric time to human friendly text
func TimeToWords(hour int, min int) string {

	nums := [...]string{"twelve", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen", "twenty", "twenty one",
		"twenty two", "twenty three", "twenty four",
		"twenty five", "twenty six", "twenty seven",
		"twenty eight", "twenty nine",
	}
	result := ""
	switch {
	case min == 0:
		result = fmt.Sprintf("%s o'clock", FirstUpper(nums[hour%12]))
	case min == 1:
		result = fmt.Sprintf("One past %s", nums[hour%12])
	case min == 59:
		result = fmt.Sprintf("One to %s", nums[(hour%12)+1])
	case min == 15:
		result = fmt.Sprintf("Quarter past %s", nums[hour%12])
	case min == 30:
		result = fmt.Sprintf("Half past %s", nums[hour%12])
	case min == 45:
		result = fmt.Sprintf("Quarter to %s", nums[(hour%12)+1])
	case min < 30:
		result = fmt.Sprintf("%s past %s", FirstUpper(nums[min]), nums[hour%12])
	case min > 30:
		result = fmt.Sprintf("%s to %s", FirstUpper(nums[60-min]), nums[(hour%12)+1])
	}
	return result
}

//This function Upper first letter
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
