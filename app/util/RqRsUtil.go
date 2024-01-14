package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func ValidateUpdateId(id uint) error {
	if id <= 0 {
		return errors.New("field 'id' is required for update")
	}
	return nil
}

func ValidateFileName(fileName string) (string, error) {
	extension := filepath.Ext(fileName)
	if len(extension) <= 0 {
		return "", errors.New("invalid file name; extension missing")
	}

	validExtensions := []string{".txt", ".json", ".pdf", ".zip", ".png", ".jpg", ".jpeg", ".gif"}
	if !slices.Contains(validExtensions, strings.ToLower(extension)) {
		return "", errors.New(fmt.Sprint("invalid file type; acceptable type: ", strings.Join(validExtensions, ", ")))
	}
	return extension, nil
}

func ValidateUploadAndGenerateName(file *multipart.FileHeader) (string, int, error) {
	extension, err := ValidateFileName(file.Filename)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	uploadTime := time.Now().Format("20060102150405.000")
	savedFileName := "file-" + uploadTime[:len(uploadTime)-4] + uploadTime[len(uploadTime)-3:] + extension
	return savedFileName, 0, nil
}

func EncodeForActivationLink(username string, otp string) string {
	strToEncode := username + "-" + otp
	return url.QueryEscape(base64.URLEncoding.EncodeToString([]byte(strToEncode)))
}

func DecodeFromActivationLink(encodedString string) (string, string, error) {
	queryUnescaped, err := url.QueryUnescape(encodedString)
	if err != nil {
		return "", "", errors.New(fmt.Sprint("error validate activation link, reason: ", err.Error()))
	}
	decodedByte, err := base64.URLEncoding.DecodeString(queryUnescaped)
	if err != nil {
		return "", "", errors.New(fmt.Sprint("error validate activation link, reason: ", err.Error()))
	}
	result := strings.Split(string(decodedByte), "-")
	return result[0], result[1], nil
}

func SetErrorResponse(ctx *gin.Context, err error, code int) {
	ctx.AbortWithStatusJSON(code, response.ApiResponse{
		Code:    code,
		Message: err.Error(),
	})
}

func SetSuccessResponseNoData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
	})
}

func SetSuccessResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, response.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func ShowErrorResponsePage(ctx *gin.Context, err error, code int) {
	info := gin.H{
		"title":   fmt.Sprint("HTTP Error - ", code),
		"header":  "An Error Has Occured! Sorry :(",
		"message": err.Error(),
		"year":    time.Now().Year(),
	}
	ctx.HTML(code, "error", info)
}

func ShowRegistrationResponsePage(ctx *gin.Context, data *rsdata.RegistrationData) {
	info := gin.H{
		"title":  "User Registration",
		"header": "Your Account Has Been Created :)",
		"data":   data,
		"year":   time.Now().Year(),
	}
	ctx.HTML(http.StatusOK, "registration", info)
}

func ShowActivationResponsePage(ctx *gin.Context, msg string) {
	info := gin.H{
		"title":   "User Activation",
		"header":  "You're Good To Go :)",
		"message": msg,
		"year":    time.Now().Year(),
	}
	ctx.HTML(http.StatusOK, "activation", info)
}

func ShowLoginResponsePage(ctx *gin.Context, data *rsdata.LoginData) {
	info := gin.H{
		"title":  "User Login",
		"header": "Welcome! Here's Your Login Info :)",
		"data":   data,
		"year":   time.Now().Year(),
	}
	ctx.HTML(http.StatusOK, "login", info)
}
