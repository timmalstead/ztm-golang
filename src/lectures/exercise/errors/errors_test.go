package timeparse

import "testing"

type ErrorTest struct {
	timeString string
	shouldPass bool
}

var testCases = []ErrorTest{
	{"19:00:12", true},
	{"1:3:44", true},
	{"bad", false},
	{"1:-3:44", false},
	{"0:59:59", true},
	{"", false},
	{"11:22", false},
	{"aa:bb:cc", false},
	{"5:23:", false},
}

func TestParseTimeStringErrors(t *testing.T) {
	for _, testCase := range testCases {
		var _, err = ParseTimeString(testCase.timeString)

		var testIsFailing = err != nil
		var throwingErrorOnCaseExpectedToPass = testCase.shouldPass && testIsFailing

		if throwingErrorOnCaseExpectedToPass {
			t.Errorf("%v: %v, error should be nil", testCase.timeString, err)
		}
	}
}
