package utils

import (
	"crypto/md5"
	"encoding/base64"
	"io"
)

func GetMd5FromFile(file io.ReadCloser) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil

}
