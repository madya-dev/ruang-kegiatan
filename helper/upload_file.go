package helper

import (
	"madyasantosa/ruangkegiatan/pkg/s3"
	"mime/multipart"
)

func UploadToS3(file *multipart.FileHeader, username string, startDate string) (string, error) {

	if file == nil {
		return "", nil
	}

	res, err := s3.UploadS3(file, username, startDate)

	return res, err
}
