package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func ListObject(bucket string) ([]*s3.Object, error) {
	session, err := CreateSession()
	if err != nil {
		return nil, err
	}
	client := s3.New(session)
	objectsV2Output, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}
	return objectsV2Output.Contents, nil
}

func UploadObject(bucket string, file *os.File, fileName string) error {
	session, err := CreateSession()
	if err != nil {
		return err
	}
	uploader := s3manager.NewUploader(session)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return err
	}

	client := s3.New(session)
	return client.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
}

func DownloadObject(bucket string, file *os.File, fileName string) (int64, error) {
	session, err := CreateSession()
	if err != nil {
		return 0, err
	}
	downloader := s3manager.NewDownloader(session)
	return downloader.Download(file, &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(fileName)})
}

func CopyObject(sourceBucket string, destBucket string, fileName string) error {
	session, err := CreateSession()
	if err != nil {
		return err
	}
	client := s3.New(session)
	_, err = client.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(destBucket),
		CopySource: aws.String(sourceBucket + "/" + fileName),
		Key:        aws.String(fileName),
	})
	if err != nil {
		return err
	}
	return client.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(destBucket),
		Key:    aws.String(fileName),
	})
}

func DeleteObject(bucket string, fileName string) error {
	session, err := CreateSession()
	if err != nil {
		return err
	}
	client := s3.New(session)
	_, err = client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return err
	}
	return client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
}
