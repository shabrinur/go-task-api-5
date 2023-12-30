package rsdata

import "idstar-idp/rest-api/app/dto"

type OtpBackupData struct {
	Info string `json:"info"`
	dto.OtpInfo
	Error string `json:"error,omitempty"`
}

type LoginData struct {
	dto.UserInfo
	dto.TokenInfo
}

type RegistrationData struct {
	OtpBackupData
	dto.UserInfo
}
