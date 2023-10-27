package s3

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"

	"madyasantosa/ruangkegiatan/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var uploader *s3manager.Uploader

func NewUploader(c config.Config) {
	s3Config := &aws.Config{
		Region:      aws.String(c.S3Region),
		Credentials: credentials.NewStaticCredentials(c.S3SAccessKey, c.S3SecretKey, ""),
	}

	s3Session := session.Must(session.NewSession(s3Config))

	uploader = s3manager.NewUploader(s3Session) // Set the package-level uploader
}

func UploadS3(fileHeader *multipart.FileHeader, username string, startDate string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return "", err
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file content: %v\n", err)
		return "", err
	}

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String("ruang-kegiatan"),                   // bucket's name
		Key:         aws.String("document/" + username + startDate), // file's destination location
		Body:        bytes.NewReader(fileContent),                   // content of the file
		ContentType: aws.String("application/pdf"),                  // content type
		ACL:         aws.String("public-read"),
	}

	res, err := uploader.UploadWithContext(context.Background(), upInput)
	if err != nil {
		log.Printf("Error uploading to S3: %v\n", err)
		return "", err
	}

	return res.Location, nil
}

func DeleteFileS3(username string, startDate string) error {

	_, err := uploader.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("ruang-kegiatan"),                   // bucket's name
		Key:    aws.String("document/" + username + startDate), // file's destination location
	})

	if err != nil {
		log.Printf("Error delete file in S3: %v\n", err)
		return err
	}
	return nil
}
