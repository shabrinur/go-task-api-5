package util

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(tokenTTL time.Duration, payload dto.AuthTokenPayload, issuer string, audiences string, jwtKey string) (*dto.TokenInfo, error) {
	iat := jwt.NewNumericDate(time.Now())
	exp := jwt.NewNumericDate(time.Now().Add(tokenTTL))

	registeredClaims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Audience:  strings.Split(audiences, ","),
		IssuedAt:  iat,
		ExpiresAt: exp,
	}

	payload.RegisteredClaims = *registeredClaims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, errors.New(fmt.Sprint("error on generating token: ", err.Error()))
	}

	tokenInfo := &dto.TokenInfo{
		AuthToken: tokenString,
		IssuedAt:  time.Unix(iat.Unix(), 0),
		ExpiredOn: time.Unix(exp.Unix(), 0),
	}

	return tokenInfo, nil
}

func ValidateToken(tokenString string, jwtKey string) (*dto.AuthTokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.AuthTokenPayload{}, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprint("error on parsing token: ", err.Error()))
	}

	tokenPayload, ok := token.Claims.(*dto.AuthTokenPayload)
	if !ok || !token.Valid {
		return tokenPayload, errors.New("token invalid; re-login to generate new token")
	}

	return tokenPayload, nil
}

func CheckPermission(tokenPayload *dto.AuthTokenPayload, path string, method string) bool {
	if strings.EqualFold(tokenPayload.RoleType, "superadmin") {
		return true
	} else {
		permissions := tokenPayload.Permissions
		for _, permission := range permissions {
			if strings.HasPrefix(path, permission.Path) {
				switch method {
				case "DELETE":
					return permission.DeleteAllowed
				case "PUT":
					return permission.PutAllowed
				case "POST":
					return permission.PostAllowed
				case "GET":
					return permission.GetAllowed
				}
			}
		}
	}
	return false
}
