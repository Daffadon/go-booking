package utils

import (
	"fmt"
	"go-booking/dto"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ImageFileExtChecker(filename string) bool {
	re := regexp.MustCompile(`(?i)\.(jpeg|jpg|png)$`)
	return re.MatchString(filename)
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}

var WORKDIR = "assets"

func UploadFile(file *multipart.FileHeader, path string) error {
	parts := strings.Split(path, "/")
	fileId := parts[1]
	dirPath := fmt.Sprintf("%s/%s", WORKDIR, parts[0])

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0750); err != nil {
			return dto.ErrFailedToCreateBook
		}
	}
	fullFilePath := fmt.Sprintf("%s/%s", dirPath, fileId)
	filepath.Clean(fullFilePath)
	uploadedFile, err := file.Open()
	if err != nil {
		return dto.ErrFailedToCreateBook
	}
	defer uploadedFile.Close()
	targetFile, err := os.Create(fullFilePath)
	if err != nil {
		return dto.ErrFailedToCreateBook
	}

	defer targetFile.Close()
	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		return dto.ErrFailedToCreateBook
	}
	return nil
}

func DeleteFile(path string) error {
	fullPath := fmt.Sprintf("%s/%s", WORKDIR, path)
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}
	return nil
}
