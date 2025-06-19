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
	title     string
	validator func(int) bool
}

func ParseTimeString(s string) (Time, error) {
	var splitTimeString = strings.Split(s, ":")
	var timeStrLen = len(splitTimeString)

	if timeStrLen != 3 {
		return Time{}, fmt.Errorf("incorrect number of time units returned from time string: %v", timeStrLen)
	}

	var timeUnits = []TimeInfo{
		{"hour", func(hour int) bool { return hour > 23 || hour < 0 }},
		{"minute", func(minute int) bool { return minute > 59 || minute < 0 }},
		{"second", func(second int) bool { return second > 59 || second < 0 }}}
	var intArr = []int{}
	for i, timing := range splitTimeString {
		// short for ascii to int. just a terrible name
		var convertedInt, err = strconv.Atoi(timing)

		if err != nil {
			return Time{}, fmt.Errorf("error parsing %v: %v", timeUnits[i].title, err)
		} else {
			intArr = append(intArr, convertedInt)
		}
	}

	for i, timeInt := range intArr {
		var timeIsOutOfBounds = timeUnits[i].validator(timeInt)

		if timeIsOutOfBounds {
			return Time{}, fmt.Errorf("%v out of range: %v", timeUnits[i].title, timeInt)
		}
	}

	var hour, minute, second = intArr[0], intArr[1], intArr[2]
	return Time{hour, minute, second}, nil
}
