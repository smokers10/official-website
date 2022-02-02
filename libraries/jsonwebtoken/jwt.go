package jsonwebtoken

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	config "github.com/smokers10/official-website.git/libraries"
)

type JwonWebToken struct {
	Payload Payload
}

type Payload struct {
	ID       int
	Email    string
	IsLogged bool
}

func Init(payload *Payload) *JwonWebToken {
	return &JwonWebToken{Payload: *payload}
}

func (j *JwonWebToken) Sign() (string, error) {
	config := config.ReadConfig().Application

	claims := jwt.MapClaims{
		"ID":       j.Payload.ID,
		"Email":    j.Payload.Email,
		"IsLogged": j.Payload.IsLogged,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (j *JwonWebToken) Verificate(tokenString string) (*Payload, error) {
	payload := j.Payload
	config := config.ReadConfig().Application

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payload.ID = int(claims["ID"].(float64))
		payload.Email = claims["Email"].(string)
		payload.IsLogged = claims["IsLogged"].(bool)

		return &payload, nil
	}

	return nil, err
}
