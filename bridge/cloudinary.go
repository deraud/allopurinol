package bridge

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryBridgeItf interface {
	UploadImage(asset interface{}) (string, error)
}

type CloudinaryBridgeImpl struct {
	cld *cloudinary.Cloudinary
	ctx context.Context
}

func NewCloudinaryBridge(cld *cloudinary.Cloudinary, ctx context.Context) CloudinaryBridgeItf {
	return &CloudinaryBridgeImpl{cld: cld, ctx: ctx}
}

func (r *CloudinaryBridgeImpl) UploadImage(asset interface{}) (string, error) {
	res, err := r.cld.Upload.Upload(r.ctx, asset, uploader.UploadParams{})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil
}
