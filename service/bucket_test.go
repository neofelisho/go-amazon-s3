package service

import (
	"github.com/google/uuid"
	"github.com/neofelisho/go-amazon-s3/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucket(t *testing.T) {
	bucketName := uuid.New().String()

	client := getS3Client(t, true)

	err := CreateBucket(client, bucketName)
	assert.NoError(t, err)

	buckets, err := ListBuckets(client)
	assert.NoError(t, err)
	namesOfBucket := helper.GetNamesOfBucket(buckets)
	assert.Contains(t, namesOfBucket, bucketName)

	err = DeleteBucket(client, bucketName)
	assert.NoError(t, err)

	buckets, err = ListBuckets(client)
	assert.NoError(t, err)
	namesOfBucket = helper.GetNamesOfBucket(buckets)
	assert.NotContains(t, namesOfBucket, bucketName)
}

func getS3Client(t *testing.T, isMock bool) s3Client {
	if isMock {
		return MockS3Client{mockS3: make(map[string]map[string]*mockS3Object)}
	}
	session, err := CreateSession()
	assert.NoError(t, err)
	return S3Client{session: session}
}
