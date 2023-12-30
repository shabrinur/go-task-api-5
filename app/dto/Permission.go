package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserPermission struct {
	Role        string       `json:"role"`
	Permissions []Permission `json:"permissions"`
}

type UserInfo struct {
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	IsUserActive bool   `json:"isUserActive"`
	UserPermission
}

type OtpInfo struct {
	Otp       string    `json:"otp,omitempty"`
	ExpiredOn time.Time `json:"expiredOn,omitempty"`
}

type TokenInfo struct {
	AuthToken string    `json:"authToken,omitempty"`
	TokenType string    `json:"tokenType,omitempty"`
	IssuedAt  time.Time `json:"issuedAt,omitempty"`
	ExpiredOn time.Time `json:"expiredOn,omitempty"`
}

type AuthTokenPayload struct {
	Username    string       `json:"username"`
	AuthMethod  string       `json:"authMethod"`
	RoleType    string       `json:"roleType"`
	Permissions []Permission `json:"permissions"`
	jwt.RegisteredClaims
}

type Permission struct {
	Path          string `json:"basePath"`
	GetAllowed    bool   `json:"getAllowed"`
	PutAllowed    bool   `json:"putAllowed"`
	PostAllowed   bool   `json:"postAllowed"`
	DeleteAllowed bool   `json:"deleteAllowed"`
}
