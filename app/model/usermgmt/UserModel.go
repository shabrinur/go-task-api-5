package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID                      uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate             time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate             *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate             gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	Fullname                string         `json:"fullname"`
	Username                string         `json:"username" gorm:"type:varchar(100);required;unique;not null"`
	Oauth                   sql.NullBool   `json:"oauth" gorm:"default:false"`
	OauthProvider           string         `json:"oauthProvider,omitempty" gorm:"type:varchar(50)"`
	Password                string         `json:"-" gorm:"type:varchar(100)"`
	AccessToken             string         `json:"-"`
	AccessTokenExpiredDate  time.Time      `json:"accessTokenExpiredDate,omitempty"`
	RefreshToken            string         `json:"-"`
	RefreshTokenExpiredDate time.Time      `json:"refreshTokenExpiredDate,omitempty"`
	Otp                     string         `json:"otp" gorm:"type:varchar(10);required;not null"`
	OtpExpiredDate          time.Time      `json:"otpExpiredDate" gorm:"required;not null"`
	AccountActivated        sql.NullBool   `json:"accountActivated" gorm:"default:false"`
	IDRole                  uint           `json:"-"`
	Role                    RoleModel      `json:"-" gorm:"foreignKey:IDRole;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (c *UserModel) TableName() string {
	return "oauth_user"
}
