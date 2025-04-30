package service

import (
	"context"
	"healthcare-allopurinol/bridge"
	"mime/multipart"
)

type ImageServiceItf interface {
	UploadImage(ctx context.Context, image *multipart.File) (string, error)
}

type ImageServiceImpl struct {
	cloudinaryBridge bridge.CloudinaryBridgeItf
}

func NewImageService(cloudinaryBridge bridge.CloudinaryBridgeItf) ImageServiceItf {
	return &ImageServiceImpl{
		cloudinaryBridge: cloudinaryBridge,
	}
}

func (s *ImageServiceImpl) UploadImage(ctx context.Context, file *multipart.File) (string, error) {
	url, err := s.cloudinaryBridge.UploadImage(*file)
	if err != nil {
		return "", err
	}
	return url, nil
}
