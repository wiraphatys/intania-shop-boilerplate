package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
)

func JwtParseToken(reqToken, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func JwtSignAccessToken(userID, secretKey string, expiration int) (*string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(time.Second * time.Duration(expiration)).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  config.GetConfig().GetServer().Name,
		"aud":  config.GetConfig().GetServer().Name,
		"type": "access",
	})

	accessTokenString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &accessTokenString, nil
}
