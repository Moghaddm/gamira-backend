package configs

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() *minio.Client {
	address := os.Getenv("MINIO_ADDRESS")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	opt := minio.Options{Creds: credentials.NewStaticV4(accessKey, secretKey, "")}
	client, err := minio.New(address, &opt)
	if err != nil {
		panic(err)
	}

	return client
}
