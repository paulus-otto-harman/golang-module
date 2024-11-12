package gola

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/paulus-otto-harman/golang-module/config"
	"github.com/paulus-otto-harman/golang-module/web/collections"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func StoreUploadedFile(inputName string, mandatory bool, r *http.Request, customUploadDir ...string) collections.FileUpload {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return collections.FileUpload{
			Error: &collections.Error{
				Default: err,
				Code:    http.StatusUnprocessableEntity,
				Message: "File size too large (Max 10MB)",
			},
		}
	}

	file, fileHandler, err := r.FormFile(inputName)
	if err != nil && mandatory {
		return collections.FileUpload{
			Error: &collections.Error{
				Default: err,
				Code:    http.StatusUnprocessableEntity,
				Message: fmt.Sprintf("%s is required", inputName),
			},
		}
	}
	if err != nil {
		return collections.FileUpload{
			Error: &collections.Error{
				Default: err,
				Code:    0,
				Message: "",
			},
		}
	}
	defer file.Close()
	fileExtension := fileHandler.Filename[strings.LastIndex(fileHandler.Filename, "."):]

	uploadDir := config.UploadDir
	if len(customUploadDir) > 0 {
		uploadDir = customUploadDir[0]
	}

	err = nil
	if uploadDirExists, _ := folderExists(uploadDir); !uploadDirExists {
		err = os.Mkdir(uploadDir, 0777)
	}

	if err != nil {
		return collections.FileUpload{Error: &collections.Error{
			Default: err,
			Code:    0,
			Message: "Unable to create upload Folder",
		}}
	}

	fileRenamed := filepath.Join(uploadDir, uuid.New().String()+fileExtension)
	destination, err := os.Create(fileRenamed)
	if err != nil {
		return collections.FileUpload{
			Error: &collections.Error{
				Default: err,
				Code:    http.StatusInternalServerError,
				Message: "Invalid server path",
			},
		}
	}
	defer destination.Close()

	if _, err = io.Copy(destination, file); err != nil {
		return collections.FileUpload{
			Error: &collections.Error{
				Default: err,
				Code:    http.StatusInternalServerError,
				Message: "Unable to store file at server",
			},
		}
	}

	scheme := "https"
	if r.TLS == nil {
		scheme = "http"
	}

	return collections.FileUpload{
		Error: nil,
		Uploaded: collections.UploadedFile{
			Path:         fileRenamed,
			MimeType:     fileHandler.Header,
			Size:         fileHandler.Size,
			OriginalName: fileHandler.Filename,
			FullUrl:      strings.Join([]string{scheme, "://", r.Host, "/", fileRenamed}, ""),
		},
	}
}

func folderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
