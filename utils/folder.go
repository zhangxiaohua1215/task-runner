package utils

import (
	"path/filepath"
	"time"
)

func GenPath(fileName string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join("upload", folderName, fileName)
	//使用mkdirall会创建多层级目录
	// os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}
