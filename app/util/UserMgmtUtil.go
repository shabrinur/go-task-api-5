package util

type UserMgmtUtil struct {
	OtpUtil  OtpUtil
	MailUtil MailUtil
}

func InitUserMgmtUtil() *UserMgmtUtil {
	return &UserMgmtUtil{
		OtpUtil:  *NewOtpUtil(),
		MailUtil: *NewMailUtil(),
	}
}
