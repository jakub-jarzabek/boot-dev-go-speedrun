package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateAccessToken(id int, secret string) (string, error) {
	var claims jwt.RegisteredClaims

	claims = jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		Subject:   fmt.Sprintf("%d", id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err

	}
	return signedJwt, nil
}

func generateRefreshToken(id int, secret string) (string, error) {
	var claims jwt.RegisteredClaims

	claims = jwt.RegisteredClaims{
		Issuer:    "chirpy-refresh",
		Subject:   fmt.Sprintf("%d", id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 60)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err

	}
	return signedJwt, nil
}

func validateAccessToken(token string, secret string) (id int, err error) {
	claims := jwt.RegisteredClaims{}

	_, err = jwt.ParseWithClaims(token, &claims, func(x *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims.Issuer != "chirpy-access" {
		return 0, errors.New("Invalid token")
	}

	subject := claims.Subject

	id, err = strconv.Atoi(subject)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func validateRefreshToken(token string, secret string) (id int, err error) {
	claims := jwt.RegisteredClaims{}

	_, err = jwt.ParseWithClaims(token, &claims, func(x *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims.Issuer != "chirpy-refresh" {
		return 0, errors.New("Invalid token")
	}

	subject := claims.Subject

	id, err = strconv.Atoi(subject)

	if err != nil {
		return 0, err
	}

	return id, nil
}
