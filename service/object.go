package service

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func ListObjects(client s3Client, bucket string) ([]*s3.Object, error) {
	return client.listObjects(bucket)
}

func UploadObject(client s3Client, bucket string, file *os.File, fileName string) error {
	return client.uploadObject(bucket, fileName, file)
}

func DownloadObject(client s3Client, bucket string, file *os.File, fileName string) (int64, error) {
	return client.downloadObject(bucket, fileName, file)
}

func CopyObject(client s3Client, sourceBucket string, destBucket string, fileName string) error {
	return client.copyObject(sourceBucket, destBucket, fileName)
}

func DeleteObject(client s3Client, bucket string, fileName string) error {
	return client.deleteObject(bucket, fileName)
}
