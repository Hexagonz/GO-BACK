package utils

import (
	"crypto/rand"
	"github.com/kataras/iris/v12"
)

var (
	App = iris.New()
)


func IfElse(condition bool, trueValue, falseValue string) string {
	if condition {
		return trueValue
	}
	return falseValue
}

func GenerateSecretKey() ([]byte, error) {
	secret := make([]byte, 16)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}