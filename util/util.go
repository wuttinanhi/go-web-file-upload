package util

import (
	"os"

	"github.com/wuttinanhi/go-web-file-upload/config"
)

func CreateDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	}
}

func IsAllowedContentType(contentType string) bool {
	for _, allowedType := range config.GetConfig().ALLOWED_CONTENT_TYPE {
		if allowedType == contentType {
			return true
		}
	}
	return false
}
