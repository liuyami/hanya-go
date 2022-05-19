// Package file 文件操作辅助函数

package file

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}

	return true
}

// FileNameWithoutExtension 去掉一个文件的后缀名并返回文件名
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string

	avatarSavePath := "public/upload/avatar/"

	// 保存文件
	// fileName := auth.CurrentUID(c) + path.Ext(file.Filename)
	fileName := "1" + path.Ext(file.Filename)

	// public/uploads/avatars/2.png
	avatarPath := avatarSavePath + fileName

	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	return avatarPath, nil
}
