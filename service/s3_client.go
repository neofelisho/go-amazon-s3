package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/neofelisho/go-amazon-s3/helper"
	"io"
)

type S3Client struct {
	session *session.Session
}

type s3Client interface {
	createBucket(bucket string) error
	listBuckets() ([]*s3.Bucket, error)
	deleteBucket(bucket string) error
	listObjects(bucket string) ([]*s3.Object, error)
	uploadObject(bucket string, objectKey string, reader io.Reader) error
	downloadObject(bucket string, objectKey string, writer io.WriterAt) (int64, error)
	copyObject(sourceBucket string, destBucket string, objectKey string) error
	deleteObject(bucket string, objectKey string) error
}

func (c S3Client) createBucket(bucket string) error {
	if !helper.ValidateBucketName(bucket) {
		return fmt.Errorf("invalid bucket name: %v", bucket)
	}
	client := s3.New(c.session)
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	return client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
}

func (c S3Client) listBuckets() ([]*s3.Bucket, error) {
	client := s3.New(c.session)
	listBucketsOutput, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}
	return listBucketsOutput.Buckets, nil
}

func (c S3Client) deleteBucket(bucket string) error {
	client := s3.New(c.session)
	_, err := client.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(bucket)})
	if err != nil {
		return err
	}
	return client.WaitUntilBucketNotExists(&s3.HeadBucketInput{Bucket: aws.String(bucket)})
}

func (c S3Client) listObjects(bucket string) ([]*s3.Object, error) {
	client := s3.New(c.session)
	objectsV2Output, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}
	return objectsV2Output.Contents, nil
}

func (c S3Client) uploadObject(bucket string, objectKey string, reader io.Reader) error {
	uploader := s3manager.NewUploader(c.session)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   reader,
	})
	if err != nil {
		return err
	}

	client := s3.New(c.session)
	return client.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
}

func (c S3Client) downloadObject(bucket string, objectKey string, writer io.WriterAt) (int64, error) {
	downloader := s3manager.NewDownloader(c.session)
	return downloader.Download(writer, &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(objectKey)})
}

func (c S3Client) copyObject(sourceBucket string, destBucket string, objectKey string) error {
	client := s3.New(c.session)
	_, err := client.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(destBucket),
		CopySource: aws.String(sourceBucket + "/" + objectKey),
		Key:        aws.String(objectKey),
	})
	if err != nil {
		return err
	}
	return client.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(destBucket),
		Key:    aws.String(objectKey),
	})
}

func (c S3Client) deleteObject(bucket string, objectKey string) error {
	client := s3.New(c.session)
	_, err := client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}
	return client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
}
