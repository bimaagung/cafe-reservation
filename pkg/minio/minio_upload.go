package minioUpload

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/minio/minio-go/v7"
)

func UploadFile(file *multipart.FileHeader, bucketName, objectName string) error {
	ctx := context.Background()

	buffer, errBuffer := file.Open()

	if errBuffer != nil {
		return errBuffer
	}

	defer buffer.Close()

	minioClient, errClient := MinioConnection(bucketName)

	if errClient != nil {
		exception.Error(errClient.Error())
	}

	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	// Upload the zip file with PutObject
	info, errInfo := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if errInfo != nil {
		return errInfo
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return nil
}