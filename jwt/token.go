package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Create token for UserClaim
func CreateUserToken(userclaim *UserClaims, secretKey string) (string, error) {
	userclaim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaim)
	serect, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return serect, nil
}

// Parse token to UserClaim
func ParseUserToken(t string, secretKey string) (userclaim *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(t, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return userclaim, err
	}
	if userclaim, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return userclaim, nil
	} else {
		return nil, err
	}
}

func CreateAuthenticateToken(claim *AuthenticateClaims, secretKey string) (string, error) {
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	serect, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return serect, nil
}

func ParseAuthenticateToken(t string, secretKey string) (claim *AuthenticateClaims, err error) {
	token, err := jwt.ParseWithClaims(t, &AuthenticateClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*AuthenticateClaims); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}
