package minio

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/minio/minio-go/v7"
)

func NewMinioRepository(minioConnection *minio.Client) domain.MinioRepository{ 
	return &minioRepositoryImpl{
		MinioConnection: minioConnection,
	}
}

type minioRepositoryImpl struct {
	MinioConnection *minio.Client
}

func (mc *minioRepositoryImpl)Upload(file *multipart.FileHeader, bucketName string, objectName string) error {
	ctx := context.Background()

	buffer, errBuffer := file.Open()

	if errBuffer != nil {
		return errBuffer
	}

	defer buffer.Close()
	
	// Create a bucket at region 'us-east-1' with object locking enabled.
	err := mc.MinioConnection.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: os.Getenv("MINIO_REGION")})
	if err != nil {
		 exists, errBucketExists := mc.MinioConnection.BucketExists(ctx, bucketName)
        if errBucketExists == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalln(err)
        }

	} else {
        log.Printf("Successfully created %s\n", bucketName)
    }

	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	// Upload the zip file with PutObject
	info, errInfo := mc.MinioConnection.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if errInfo != nil {
		return errInfo
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return nil
}