package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/neofelisho/go-amazon-s3/helper"
)

func CreateBucket(bucket string) error {
	if !helper.ValidateBucketName(bucket) {
		return fmt.Errorf("invalid bucket name: %v", bucket)
	}
	session, err := CreateSession()
	if err != nil {
		return err
	}
	client := s3.New(session)
	_, err = client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	return client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
}

func ListBuckets() ([]*s3.Bucket, error) {
	session, err := CreateSession()
	if err != nil {
		return nil, err
	}
	client := s3.New(session)
	listBucketsOutput, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}
	return listBucketsOutput.Buckets, nil
}

func DeleteBucket(bucket string) error {
	session, err := CreateSession()
	if err != nil {
		return err
	}
	client := s3.New(session)
	_, err = client.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(bucket)})
	if err != nil {
		return err
	}
	return client.WaitUntilBucketNotExists(&s3.HeadBucketInput{Bucket: aws.String(bucket)})
}
