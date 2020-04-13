package service

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/neofelisho/go-amazon-s3/helper"
	"io"
)

type MockS3Client struct {
	// The first map key is bucket name, and the second one is object key.
	mockS3 map[string]map[string]*mockS3Object
}

type mockS3Object struct {
	size    int64
	content []byte
}

func (c MockS3Client) createBucket(bucket string) error {
	if !helper.ValidateBucketName(bucket) {
		return fmt.Errorf("invalid bucket name: %v", bucket)
	}
	if c.mockS3[bucket] != nil {
		return fmt.Errorf("existing bucket name: %v", bucket)
	}
	c.mockS3[bucket] = make(map[string]*mockS3Object)
	return nil
}

func (c MockS3Client) listBuckets() ([]*s3.Bucket, error) {
	buckets := make([]*s3.Bucket, 0)
	for b := range c.mockS3 {
		bucket := s3.Bucket{Name: aws.String(b)}
		buckets = append(buckets, &bucket)
	}
	return buckets, nil
}

func (c MockS3Client) deleteBucket(bucket string) error {
	if c.mockS3[bucket] != nil && len(c.mockS3[bucket]) > 0 {
		return fmt.Errorf("can't delete non-empty bucket: %v", bucket)
	}
	delete(c.mockS3, bucket)
	return nil
}

func (c MockS3Client) listObjects(bucket string) ([]*s3.Object, error) {
	if c.mockS3[bucket] == nil {
		return nil, fmt.Errorf("bucket doesn't exist: %v", bucket)
	}
	objects := make([]*s3.Object, 0)
	for key := range c.mockS3[bucket] {
		object := s3.Object{Key: aws.String(key)}
		objects = append(objects, &object)
	}
	return objects, nil
}

func (c MockS3Client) uploadObject(bucket string, objectKey string, reader io.Reader) error {
	if c.mockS3[bucket] == nil {
		return fmt.Errorf("bucket doesn't exist: %v", bucket)
	}
	buffer := new(bytes.Buffer)
	n, err := buffer.ReadFrom(reader)
	if err != nil {
		return err
	}
	c.mockS3[bucket][objectKey] = &mockS3Object{
		size:    n,
		content: buffer.Bytes(),
	}
	return nil
}

func (c MockS3Client) downloadObject(bucket string, objectKey string, writer io.WriterAt) (int64, error) {
	if c.mockS3[bucket][objectKey] == nil {
		return 0, fmt.Errorf("object doesn't exist: %v/%v", bucket, objectKey)
	}
	n, err := writer.WriteAt(c.mockS3[bucket][objectKey].content, 0)
	return int64(n), err
}

func (c MockS3Client) copyObject(sourceBucket string, destBucket string, objectKey string) error {
	if c.mockS3[sourceBucket] == nil {
		return fmt.Errorf("bucket doesn't exist: %v", sourceBucket)
	}
	if c.mockS3[destBucket] == nil {
		return fmt.Errorf("bucket doesn't exist: %v", destBucket)
	}
	if c.mockS3[sourceBucket][objectKey] == nil {
		return fmt.Errorf("object doesn't exist: %v/%v", sourceBucket, objectKey)
	}
	c.mockS3[destBucket][objectKey] = c.mockS3[sourceBucket][objectKey]
	return nil
}

func (c MockS3Client) deleteObject(bucket string, objectKey string) error {
	if c.mockS3[bucket] == nil {
		return fmt.Errorf("bucket doesn't exist: %v", bucket)
	}
	if c.mockS3[bucket][objectKey] == nil {
		return fmt.Errorf("object doesn't exist: %v/%v", bucket, objectKey)
	}
	delete(c.mockS3[bucket], objectKey)
	return nil
}
