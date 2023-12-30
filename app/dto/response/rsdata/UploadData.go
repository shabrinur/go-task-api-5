package rsdata

type UploadData struct {
	FileName        string `json:"fileName"`
	FileDownloadUri string `json:"fileDownloadUri"`
	FileType        string `json:"fileType"`
	Size            int64  `json:"size"`
}
