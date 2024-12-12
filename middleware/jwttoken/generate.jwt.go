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

func GenerateTokenJwt(email string, id string, ctx iris.Context) (*ResponseGenerate, error) {
	accessClaims := Claims{
		Claims: jwt.Claims{Subject: email},
		ID:     id,
		Email:  email,
	}
	accsesToken, err := AccessSigner.Sign(accessClaims)
	if err != nil {
		fmt.Println(err)
		return &ResponseGenerate{}, err
	}

	refreshClaims := Claims{
		Claims: jwt.Claims{Subject: email},
		ID:     id,
		Email:  email,
	}
	refrshToken, err := RefreshSigner.Sign(refreshClaims)
	if err != nil {
		return &ResponseGenerate{}, err
	}

	return &ResponseGenerate{
		AccessToken:      string(accsesToken),
		ExpiresAt:        int64(accessExpire.Seconds()),
		RefreshToken:     string(refrshToken),
		RefreshExpiresAt: int64(refreshExpire.Seconds()),
	}, nil

}
