package service

import (
	"github.com/google/uuid"
	"github.com/neofelisho/go-amazon-s3/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucket(t *testing.T) {
	bucketName := uuid.New().String()

	err := CreateBucket(bucketName)
	assert.NoError(t, err)

	buckets, err := ListBucket()
	assert.NoError(t, err)
	namesOfBucket := helper.GetNamesOfBucket(buckets)
	assert.Contains(t, namesOfBucket, bucketName)

	err = DeleteBucket(bucketName)
	assert.NoError(t, err)

	buckets, err = ListBucket()
	assert.NoError(t, err)
	namesOfBucket = helper.GetNamesOfBucket(buckets)
	assert.NotContains(t, namesOfBucket, bucketName)
}
