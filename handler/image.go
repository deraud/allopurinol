package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandlerItf interface {
	PostHandler(ctx *gin.Context)
}

type ImageHandlerImpl struct {
	imageService service.ImageServiceItf
}

func NewImageHandler(imageService service.ImageServiceItf) ImageHandlerItf {
	return &ImageHandlerImpl{
		imageService: imageService,
	}
}

func (h *ImageHandlerImpl) PostHandler(ctx *gin.Context) {
	var body dto.UploadImageRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	file, err := body.Image.Open()
	if err != nil {
		ctx.Error(err)
		return
	}

	url, err := h.imageService.UploadImage(ctx, &file)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: gin.H{"url": url},
	})
}
