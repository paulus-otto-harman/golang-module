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

func StoreUploadedFile(inputName string, mandatory bool, r *http.Request) collections.FileUpload {
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

	fileRenamed := filepath.Join(config.UploadDir, uuid.New().String()+fileExtension)
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
