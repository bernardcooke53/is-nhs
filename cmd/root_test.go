package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type nhsNumberValidTestCase struct {
	Input    int
	Expected bool
}

var nhsNumberValidTestCases = []nhsNumberValidTestCase{
	{1021580554, true},
	{123456, false},
	{23333333333333, false},
	{1420002454, false},
}

func TestNHSNumberValid(t *testing.T) {
	for _, testCase := range nhsNumberValidTestCases {
		assert.Equal(t, isValidNHSNumber(testCase.Input), testCase.Expected)
	}
}

type intToDigitsTestCase struct {
	Input    int
	Expected []int
}

var intToDigitsTestCases = []intToDigitsTestCase{
	{123, []int{1, 2, 3}},
	{245856, []int{2, 4, 5, 8, 5, 6}},
}

func TestIntToDigits(t *testing.T) {
	for _, testCase := range intToDigitsTestCases {
		assert.Equal(t, intToDigits(testCase.Input), testCase.Expected)
	}
}
