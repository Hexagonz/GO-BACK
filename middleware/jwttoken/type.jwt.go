package jwttoken

import (
	"fmt"

	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	privateKey, publicKey  = jwt.MustLoadRSA("./private/private.pem", "./private/public.pem")
	priRefresh, pubRefresh = jwt.MustLoadRSA("./private/private_key.pem", "./private/public_key.pem")

	AccessSigner  = jwt.NewSigner(jwt.RS256, privateKey, accessExpire).WithEncryption(secretKey, nil)
	RefreshSigner = jwt.NewSigner(jwt.RS256, priRefresh, refreshExpire)

	Verifier        = jwt.NewVerifier(jwt.RS256, publicKey).WithDecryption(secretKey, nil)
	VerifierRefresh = jwt.NewVerifier(jwt.RS256, pubRefresh)
)

type Claims struct {
	jwt.Claims
	ID    string `json:"id"`
	Email string `json:"email"`
}

type ResponseReset struct {
	AccessToken string `json:"access_token"`
	ExpiredAt int64 `json:"expired_at"`
}

func (u *Claims) Validate() error {
	if u.ID == "" {
		return fmt.Errorf("username field is missing")
	}
	if u.Email == "" {
		return fmt.Errorf("email field is missing")
	}

	return nil
}