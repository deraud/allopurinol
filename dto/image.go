package dto

import "mime/multipart"

type UploadImageRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type UploadImageResponse struct {
	Url string `json:"url"`
}
