package utils

import (
	"fmt"
)

func IfElse(condition bool, trueValue, falseValue string) string {
	if condition {
		return trueValue
	}
	return falseValue
}

func Catch(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err)) 
	}
}