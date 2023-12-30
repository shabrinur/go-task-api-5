package util

type UserMgmtUtil struct {
	OtpUtil  OtpUtil
	MailUtil MailUtil
	PwdUtil  PasswordUtil
}

func InitUserMgmtUtil() *UserMgmtUtil {
	return &UserMgmtUtil{
		OtpUtil:  *NewOtpUtil(),
		MailUtil: *NewMailUtil(),
		PwdUtil:  PasswordUtil{},
	}
}
