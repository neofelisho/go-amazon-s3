package service

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket(client s3Client, bucket string) error {
	return client.createBucket(bucket)
}

func ListBuckets(client s3Client) ([]*s3.Bucket, error) {
	return client.listBuckets()
}

func DeleteBucket(client s3Client, bucket string) error {
	return client.deleteBucket(bucket)
}
