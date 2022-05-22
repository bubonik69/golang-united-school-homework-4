package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
var re = regexp.MustCompile("[0-9]+")
func StringSum (input string) (string, error) {
		result, err := sum(input)
	return result, err

}
func sum(input string) (string, error) {
	nums, err := findNums(input)
	if err != nil {
		return "", err
	}
	input = strings.Replace(input, " ", "", -1)
	secondInput := strings.Index(input, nums[1])
	firstVarStr := input[:secondInput-1]
	firstVar, err := strToInt(firstVarStr)
	if err != nil {
		return "", fmt.Errorf("NumError")
	}
	secondVarStr := input[secondInput-1:]
	secondVar, err := strToInt(secondVarStr)
	if err != nil {
		return "", fmt.Errorf("NumError")
	}
	return strconv.Itoa(firstVar + secondVar), nil
}

func findNums(input string) ([]string, error) {
	nums := re.FindAllString(input, -1)
	countOperands := len(nums)
	switch {
	case countOperands == 0:
		return nil, fmt.Errorf("errorEmptyInput")
	case countOperands != 2:
		return nil, fmt.Errorf("errorNotTwoOperands")
	}
	return nums, nil
}

func strToInt(input string) (int, error) {
	intVar, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return intVar, nil}