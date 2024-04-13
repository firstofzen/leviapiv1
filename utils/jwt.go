package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JwtToken struct {
}

var jwtPrivateKey = os.Getenv("JWT_PRIVATE_KEY")
var j JwtToken

func (j JwtToken) CreateAccessToken(email string) (string, error) {
	if email == "" {
		return "", errors.New("cannot get id user from claims")
	} else {
		tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":    email,
			"expireAt": time.Now().Add(1440 * time.Minute).Nanosecond(),
		})
		rs, err := tok.SignedString([]byte(jwtPrivateKey))
		if err != nil {
			return "", errors.New("can not sign key")
		}
		return rs, nil
	}
}
func (j JwtToken) TokenIsExpire(str string) (bool, error) {
	if err := j.VerifyToken(str); err != nil {
		return false, err
	} else {
		if claims, err := j.ParseToken(str); err != nil {
			return false, err
		} else {
			return claims["expireAt"].(float64) < float64(time.Now().Nanosecond()), nil
		}
	}
}
func (j JwtToken) VerifyToken(str string) error {
	_, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtPrivateKey), nil
	})
	return err
}
func (j JwtToken) ParseToken(str string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtPrivateKey), nil
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return token.Claims.(jwt.MapClaims), nil
}
