package validation

import (
	"fmt"
	"healthcare-allopurinol/constant"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateImage(sl validator.StructLevel) {
	file := sl.Current().Interface().(multipart.FileHeader)

	if file.Size > constant.IMAGE_MAX_SIZE {
		sl.ReportError(file.Size, "image", "image", "size", fmt.Sprintf("%d bytes", constant.IMAGE_MAX_SIZE))
	}

	fileExt := filepath.Ext(file.Filename)
	valid := false
	for _, ext := range constant.IMAGE_EXTENSIONS {
		if fileExt == ext {
			valid = true
			break
		}
	}
	if !valid {
		sl.ReportError(file.Filename, "image", "image", "extension", strings.Join(constant.IMAGE_EXTENSIONS, " "))
	}
}
