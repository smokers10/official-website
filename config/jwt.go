package config

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JsonWebToken struct {
	ID       int
	Email    string
	IsLogged bool
}

func (j *JsonWebToken) Sign() (string, error) {
	config := ReadConfig().Application

	claims := jwt.MapClaims{
		"ID":       j.ID,
		"Email":    j.Email,
		"IsLogged": j.IsLogged,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (j *JsonWebToken) Verificate(tokenString string) (*JsonWebToken, error) {
	config := ReadConfig().Application

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
		j.ID = int(claims["ID"].(float64))
		j.Email = claims["Email"].(string)
		j.IsLogged = claims["IsLogged"].(bool)

		return j, nil
	}

	return nil, err
}
