package string_sum

import (
	"errors"
	"fmt"

	"strconv"

	//"fmt"
	//	"strconv"

	//"strconv"
	"strings"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	//"strings"
	//	"strconv"
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

func StringSum(input string) (output string, err error) {

	if input == "" {return "", fmt.Errorf("%w", errorEmptyInput)}
	
	var minus_count int
	s1 :=  strings.ReplaceAll(input, " ", "")
	for _, v := range s1 {
		if v == 45 {
			minus_count++
		}
	}

	s2 := strings.FieldsFunc(s1, Split)
	if len(s2) > 2 || len(s2) <= 1 {return "",  fmt.Errorf("%w", errorNotTwoOperands)}

	opp1, err1  := strconv.Atoi(s2[0])
	opp2, err2  := strconv.Atoi(s2[1])

	if err1 != nil {
		return "", fmt.Errorf("%w", err1)
	}else if err2 != nil {
		return "", fmt.Errorf("%w", err2)
	}

	if minus_count == 2 {
		return strconv.Itoa(-opp1-opp2), err
	}else if minus_count == 1 && s1[0] == 45 && opp1 > opp2{
		return strconv.Itoa(opp2-opp1), err
	}else if minus_count == 1 && s1[0] == 45 && opp1 < opp2{
		return strconv.Itoa(opp1-opp2), err
	}else if minus_count == 1 && opp1 < opp2{
		return strconv.Itoa(opp1-opp2), err
	}else if minus_count == 1 && opp1 > opp2{
		return strconv.Itoa(opp1-opp2), err
	}else {return strconv.Itoa(opp2+opp1), err}
	
}

func Split(r rune) bool {
    return r == '+' || r == '-'
}
