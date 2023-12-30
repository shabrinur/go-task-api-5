package controller

import (
	"errors"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileUploadController struct {
	baseUrl    string
	uploadPath string
}

func NewFileUploadController() *FileUploadController {
	return &FileUploadController{
		baseUrl:    config.GetConfigValue("app.base.url"),
		uploadPath: config.GetConfigValue("file.upload.path"),
	}
}

// UploadFile godoc
//
//	@Summary	Upload File
//	@Id			UploadFile
//	@Tags		file
//	@Accept		*/*
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		file	formData	file	true	"File Upload Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/file/upload [post]
func (ctrl *FileUploadController) UploadFile(ctx *gin.Context) {
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

	err = ctx.SaveUploadedFile(file, ctrl.uploadPath+savedFileName)
	if err != nil {
		util.SetErrorResponse(ctx, errors.New("failed to save file"), http.StatusInternalServerError)
		return
	}

	uploadData := rsdata.UploadData{
		FileName:        savedFileName,
		FileDownloadUri: ctrl.baseUrl + "/v1/file/show/" + savedFileName,
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
//	@Security	ApiKeyAuth
//	@Param		filename	path		string	true	"File Name"
//	@Response	200			{file}		file
//	@Response	404			{object}	response.ApiResponse
//	@Router		/v1/file/show/{filename} [get]
func (ctrl *FileUploadController) ShowFile(ctx *gin.Context) {
	fileName := ctx.Param("filename")

	_, err := util.ValidateFileName(fileName)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile(ctrl.uploadPath + fileName)
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
//	@Security	ApiKeyAuth
//	@Param		filename	path		string	true	"File Name"
//	@Response	200			{object}	response.ApiResponse
//	@Response	500			{object}	response.ApiResponse
//	@Router		/v1/file/delete/{filename} [delete]
func (ctrl *FileUploadController) DeleteFile(ctx *gin.Context) {
	fileName := ctx.Param("filename")

	_, err := util.ValidateFileName(fileName)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	err = os.Remove(ctrl.uploadPath + fileName)
	if err != nil {
		util.SetErrorResponse(ctx, errors.New("failed to remove file"), http.StatusInternalServerError)
		return
	}
	util.SetSuccessResponseNoData(ctx)
}
