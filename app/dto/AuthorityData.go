package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserPermission struct {
	Role        string       `json:"role"`
	Permissions []Permission `json:"permissions"`
}

type OtpBackupData struct {
	Info  string `json:"info"`
	Error string `json:"error"`
	OtpInfo
}

type OtpInfo struct {
	Otp       string    `json:"otp"`
	ExpiredOn time.Time `json:"expiredOn"`
}

type LoginData struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	UserPermission
	TokenInfo
}

type TokenInfo struct {
	AuthToken string    `json:"authToken,omitempty"`
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
