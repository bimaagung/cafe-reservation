package miniodb

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection() (*minio.Client) {
	
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

	return minioClient
}
