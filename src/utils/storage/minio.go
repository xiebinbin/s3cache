package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"s3cache/src/utils"
)

var minioClient *minio.Client = nil

func initMinioClient() {
	cfg := utils.GetConfig()
	client, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyId, cfg.Minio.SecretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	minioClient = client
}
func DownloadObject(bucket string, objectName string, filePath string) error {
	err := minioClient.FGetObject(context.Background(), bucket, objectName, filePath, minio.GetObjectOptions{})
	// 文件路径 大小 时间

	return err
}
