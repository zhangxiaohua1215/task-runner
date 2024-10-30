package utils

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"mime/multipart"
)

func GetMd5FromFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err!= nil {
		return "", err
	}
	h := md5.New()
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil

}
