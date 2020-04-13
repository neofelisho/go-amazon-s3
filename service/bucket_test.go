package service

import (
	"github.com/google/uuid"
	"github.com/neofelisho/go-amazon-s3/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucket(t *testing.T) {
	bucketName := uuid.New().String()

	session, err := CreateSession()
	assert.NoError(t, err)
	client := S3Client{session: session}

	err = CreateBucket(client, bucketName)
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
