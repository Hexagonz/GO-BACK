package jwttoken

import (
	"fmt"
	"log"
	"time"

	"github.com/Hexagonz/back-end-go/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const (
	accessExpire  = 5 * time.Minute
	refreshExpire = 24 * time.Hour
)

var (
	secretKey, err = utils.GenerateSecretKey()
)

func init() {
	if err != nil {
		log.Fatal("Failed to generate secret key: " + err.Error())
	}
}


func GenerateTokenJwt(email string, id string, ctx iris.Context) (map[string]interface{}, error) {
	accessClaims := Claims{
		Claims: jwt.Claims{Subject: email},
		ID:     id,
		Email:  email,
	}
	accsesToken, err := AccessSigner.Sign(accessClaims)
	if err != nil {
		fmt.Println(err)
		return make(map[string]interface{}), err
	}

	refreshClaims := Claims{
		Claims: jwt.Claims{Subject: email},
		ID:     id,
		Email:  email,
	}
	refrshToken, err := RefreshSigner.Sign(refreshClaims)
	if err != nil {
		return make(map[string]interface{}), err
	}

	return map[string]interface{}{
		"acces_token":   string(accsesToken),
		"refresh_token": string(refrshToken),
		"expires_at":    int64(accessExpire.Seconds()),
	}, nil

}
