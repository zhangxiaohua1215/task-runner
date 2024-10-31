package utils

import (
	"path/filepath"
	"strconv"
	"time"
)

func GenScriptPath(id int64, fileName string) string {
	return genfilePath("upload", time.Now(), id, fileName)
}

// // 生成输入文件路径
// func GenInputFilePath(createTime time.Time, id int64, fileName string) string {
// 	return genfilePath("input_files", createTime, id, fileName)

// }

// // 生成输出文件路径
// func GenResultFilePath(createTime time.Time, id int64) string {
// 	return genfilePath("output_files", createTime, id, "result.txt")
// }

// 输入文件路径
func GenInputFilePath(taskID int64, fileName string) string {
	return filepath.Join("input_files", strconv.FormatInt(taskID, 16)+"-"+fileName)
}

// 生成输出文件路径
func GenResultFilePath(taskID int64) string {
	return filepath.Join("output_files", strconv.FormatInt(taskID, 16)+"-result.txt")
}


// 通用生成文件路径， 会根据时间生成文件夹， 文件夹格式为 2006/01/02
func genfilePath(rootName string, createTime time.Time, id int64, fileName string) string {
	folderName := createTime.Format("2006/01/02")
	return filepath.Join(rootName, folderName, strconv.FormatInt(id, 10)+"-"+fileName)
}