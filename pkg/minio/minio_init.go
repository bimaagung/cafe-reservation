package minioUpload

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection(bucketName string) (*minio.Client, error) {
	ctx := context.Background()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
	useSSL := false

	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if errInit != nil {
		log.Fatalln(errInit)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	location := "us-east-1"

	// Create a bucket at region 'us-east-1' with object locking enabled.
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		 exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
        if errBucketExists == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalln(err)
        }

	} else {
        log.Printf("Successfully created %s\n", bucketName)
    }
	
	return minioClient, errInit

}
