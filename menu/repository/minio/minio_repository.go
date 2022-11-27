package repository

import "mime/multipart"

type MinioRepository interface {
	Upload(file *multipart.FileHeader, bucketName string, objectName string) error
}