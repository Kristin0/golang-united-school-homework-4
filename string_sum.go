package string_sum

import (
	"errors"
	
	"strconv"

	//"fmt"
	//	"strconv"

	//"strconv"
	"strings"
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

	if input == "" {return "", errorEmptyInput}
	
	var minus_count int
	var plus_count int
	s1 :=  strings.ReplaceAll(input, " ", "")
	for _, v := range s1 {
		if v == 45 {
			minus_count++
		}else if v == 43 { 
			plus_count++
		}else if  v > 57 || v < 48 {return "", errorNotTwoOperands}
	}
	s2 := strings.FieldsFunc(s1, Split)
	opp1, _  := strconv.Atoi(s2[0])
	opp2, _  := strconv.Atoi(s2[1])
	if minus_count > 2 || plus_count > 2 {
		return "", errorNotTwoOperands
	}else if minus_count == 2 {
		
		return strconv.Itoa(-opp1-opp2), err
	}else if minus_count == 1 && s1[0] == 45 {
		return strconv.Itoa(opp2-opp1), err
	}else{
		return strconv.Itoa(opp2+opp1), err
	}


	
}

func Split(r rune) bool {
    return r == '+' || r == '-'
}
