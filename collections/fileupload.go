package collections

import "net/textproto"

type Error struct {
	Default error
	Code    int
	Message string
}

type UploadedFile struct {
	OriginalName string
	Path         string
	Size         int64
	MimeType     textproto.MIMEHeader
}

type FileUpload struct {
	Error        *Error
	UploadedFile UploadedFile
}
