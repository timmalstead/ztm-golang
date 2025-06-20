//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeInfo struct {
	title      string
	upperBound int
}

var timeUnits = []TimeInfo{
	{"hour", 23},
	{"minute", 59},
	{"second", 59},
}

func ParseTimeString(s string) (Time, error) {
	var splitTimeString = strings.Split(s, ":")
	var splitStrLen = len(splitTimeString)

	if splitStrLen != 3 {
		return Time{}, fmt.Errorf("incorrect number of time units returned from time string: %v", splitStrLen)
	}

	var intArr = []int{}
	for i, timing := range splitTimeString {
		var currentTimeUnit = timeUnits[i]
		// short for ascii to int. just a terrible name
		var convertedInt, err = strconv.Atoi(timing)

		if err != nil {
			return Time{}, fmt.Errorf("error parsing %v: %v", currentTimeUnit.title, err)
		}

		var timeIsOutOfBounds = convertedInt > currentTimeUnit.upperBound || convertedInt < 0

		if timeIsOutOfBounds {
			return Time{}, fmt.Errorf("%v out of range. current value is %v. it should be between 0 and %v", currentTimeUnit.title, convertedInt, currentTimeUnit.upperBound)
		}

		intArr = append(intArr, convertedInt)
	}

	return Time{intArr[0], intArr[1], intArr[2]}, nil
}
