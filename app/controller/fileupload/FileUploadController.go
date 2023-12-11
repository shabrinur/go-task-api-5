package controller

import (
	"errors"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// UploadFile godoc
//
//	@Summary	Upload File
//	@Id			UploadFile
//	@Tags		file
//	@Accept		*/*
//	@Produce	json
//	@Param		file	formData	file	true	"File Upload Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/file/upload [post]
func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		util.SetErrorResponse(ctx, errors.New("failed to receive file"), http.StatusBadRequest)
		return
	}

	savedFileName, code, err := util.ValidateUploadAndGenerateName(file)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	err = ctx.SaveUploadedFile(file, "/upload/"+savedFileName)
	if err != nil {
		util.SetErrorResponse(ctx, errors.New("failed to save file"), http.StatusInternalServerError)
		return
	}

	uploadData := response.UploadData{
		FileName:        savedFileName,
		FileDownloadUri: ctx.Request.Host + "/v1/idstar/file/show/" + savedFileName,
		FileType:        file.Header.Get("Content-Type"),
		Size:            file.Size,
	}
	util.SetSuccessResponse(ctx, uploadData)
}

// ShowFile godoc
//
//	@Summary	Show Uploaded File
//	@Id			ShowFile
//	@Tags		file
//	@Accept		json
//	@Produce	*/*
//	@Param		filename	path		string	true	"File Name"
//	@Response	200			{file}		file
//	@Response	404			{object}	response.ApiResponse
//	@Router		/file/show/{filename} [get]
func ShowFile(ctx *gin.Context) {
	fileName := ctx.Param("filename")

	_, err := util.ValidateFileName(fileName)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile("/upload/" + fileName)
	if err != nil || len(data) <= 0 {
		util.SetErrorResponse(ctx, errors.New("failed to find file"), http.StatusNotFound)
		return
	}

	ctx.Header("Content-Description", "Uploaded File")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Data(http.StatusAccepted, http.DetectContentType(data), data)
}

// DeleteFile godoc
//
//	@Summary	Delete Uploaded File
//	@Id			DeleteFile
//	@Tags		file
//	@Accept		json
//	@Produce	json
//	@Param		filename	path		string	true	"File Name"
//	@Response	200			{object}	response.ApiResponse
//	@Response	500			{object}	response.ApiResponse
//	@Router		/file/delete/{filename} [delete]
func DeleteFile(ctx *gin.Context) {
	fileName := ctx.Param("filename")

	_, err := util.ValidateFileName(fileName)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	err = os.Remove("/upload/" + fileName)
	if err != nil {
		util.SetErrorResponse(ctx, errors.New("failed to remove file"), http.StatusInternalServerError)
		return
	}
	util.SetSuccessResponseNoData(ctx)
}
