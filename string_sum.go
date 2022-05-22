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

func main(){
	fmt.Println(StringSum("24-55"))
}

var re = regexp.MustCompile("[0-9]+")
func StringSum (input string) (string, error) {
	//func StringSum(input string) (output string, err error) {
	// удаляем пробелы
	var sum int
	s:=strings.ReplaceAll(input," ","")
	// узнаем количество слагаемых


	// если после удаления пробелов - длина строки нуль - пустое поле
	if len(s)==0{
		return "", fmt.Errorf("1: %w", errorEmptyInput)
	}

	term:=strings.Split(s,"+") // делим строку по плюсам
	if len(term)==2{
		for _,s:=range term{
			if s!="" {
				numer, err := strconv.Atoi(strings.TrimSpace(s))
				sum += numer
				if err != nil {
					return "", err
				}}else {
				return "",errorNotTwoOperands
			}
		}
	}
	if len(term)>2{
		return "", fmt.Errorf("2: %w", errorNotTwoOperands)
	}

	if len(term)==1{
		// проводим дополнительную проверку на минусы
		var re = regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(input, -1)
		if len(nums)!=2{
			return "", errorNotTwoOperands
		}
		firstInput := strings.Index(s, nums[0])
		border := firstInput + len(nums[0])
		firstTerm,err:=strconv.Atoi(strings.TrimSpace(s[:border]))
		if err!=nil{
			return "",err
		}
		secondTermerr,err:=strconv.Atoi(strings.TrimSpace(s[border:]))
		if err!=nil{
			return "",err
		}
		sum=firstTerm+secondTermerr
	}


	return strconv.Itoa(sum),nil
}
