package jwttoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	refreshExpired time.Duration = 15 * time.Minute
)

func RefreshTokenJWT(ctx iris.Context) (*ResponseReset, error) {
	user := jwt.Get(ctx).(*Claims)

	refreshClaims := jwt.Claims{Subject: user.Email}
	accessClaims := Claims{
		ID:    user.ID,
		Email: user.Email,
	}

	tokenPair, err := AccessSigner.NewTokenPair(accessClaims, refreshClaims, refreshExpired)
	if err != nil {
		return nil, fmt.Errorf("failed to repair refresh token: %v", err)
	}
	token := strings.Trim(string(tokenPair.AccessToken), "\"")
	return &ResponseReset{
		AccessToken: token,
		ExpiredAt:   int64(refreshExpired.Seconds()),
	}, nil
}
