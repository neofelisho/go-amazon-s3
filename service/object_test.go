package service

import (
	"github.com/google/uuid"
	"github.com/neofelisho/go-amazon-s3/helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testData struct {
	t                *testing.T
	client           s3Client
	sourceBucketName string
	destBucketName   string
	testFileName     string
	testFileContent  string
	testFileSize     int64
	downloadFileName string
}

func TestObject(t *testing.T) {
	test := initTest(t, true)
	test.createTestFile()
	test.createBuckets()
	test.uploadObject()
	test.downloadObject()
	test.copyObject()
	test.deleteObjects()
	test.deleteBuckets()
	test.deleteTestFiles()
}

func initTest(t *testing.T, isMock bool) *testData {
	client := getS3Client(t, isMock)

	return &testData{
		t:                t,
		client:           client,
		sourceBucketName: "source-" + uuid.New().String(),
		destBucketName:   "dest-" + uuid.New().String(),
		testFileName:     "test-" + uuid.New().String(),
		testFileContent:  "content-" + uuid.New().String(),
		downloadFileName: "download-" + uuid.New().String(),
	}
}

func (d *testData) createTestFile() {
	file, err := os.Create(d.testFileName)
	defer closeFile(d.t, file)
	assert.NoError(d.t, err)

	n, err := file.WriteString(d.testFileContent)
	assert.NoError(d.t, err)

	err = file.Sync()
	assert.NoError(d.t, err)
	d.testFileSize = int64(n)
}

func (d testData) createBuckets() {
	err := CreateBucket(d.client, d.sourceBucketName)
	assert.NoError(d.t, err)
	err = CreateBucket(d.client, d.destBucketName)
	assert.NoError(d.t, err)
}

func (d *testData) uploadObject() {
	file, err := os.Open(d.testFileName)
	defer closeFile(d.t, file)
	assert.NoError(d.t, err)

	err = UploadObject(d.client, d.sourceBucketName, file, d.testFileName)
	assert.NoError(d.t, err)

	objects, err := ListObjects(d.client, d.sourceBucketName)
	assert.NoError(d.t, err)
	assert.Len(d.t, objects, 1)
	assert.Equal(d.t, d.testFileName, *objects[0].Key)
}

func (d *testData) downloadObject() {
	file, err := os.Create(d.downloadFileName)
	defer closeFile(d.t, file)
	assert.NoError(d.t, err)

	size, err := DownloadObject(d.client, d.sourceBucketName, file, d.testFileName)
	assert.NoError(d.t, err)

	b := make([]byte, size)
	n, err := file.Read(b)
	assert.NoError(d.t, err)
	assert.Equal(d.t, d.testFileSize, int64(n))
	assert.Equal(d.t, d.testFileContent, string(b))
}

func (d *testData) copyObject() {
	err := CopyObject(d.client, d.sourceBucketName, d.destBucketName, d.testFileName)
	assert.NoError(d.t, err)

	objects, err := ListObjects(d.client, d.destBucketName)
	assert.NoError(d.t, err)
	assert.Len(d.t, objects, 1)
	assert.Equal(d.t, d.testFileName, *objects[0].Key)
}

func (d *testData) deleteObjects() {
	err := DeleteObject(d.client, d.sourceBucketName, d.testFileName)
	assert.NoError(d.t, err)

	objects, err := ListObjects(d.client, d.sourceBucketName)
	assert.NoError(d.t, err)
	assert.Len(d.t, objects, 0)

	err = DeleteObject(d.client, d.destBucketName, d.testFileName)
	assert.NoError(d.t, err)

	objects, err = ListObjects(d.client, d.destBucketName)
	assert.NoError(d.t, err)
	assert.Len(d.t, objects, 0)
}

func (d *testData) deleteBuckets() {
	err := DeleteBucket(d.client, d.sourceBucketName)
	assert.NoError(d.t, err)
	err = DeleteBucket(d.client, d.destBucketName)
	assert.NoError(d.t, err)

	buckets, err := ListBuckets(d.client)
	assert.NoError(d.t, err)
	namesOfBucket := helper.GetNamesOfBucket(buckets)
	assert.NotContains(d.t, namesOfBucket, d.sourceBucketName)
	assert.NotContains(d.t, namesOfBucket, d.destBucketName)
}

func (d *testData) deleteTestFiles() {
	deleteFile(d.t, d.testFileName)
	deleteFile(d.t, d.downloadFileName)
}

func closeFile(t *testing.T, file *os.File) {
	err := file.Close()
	assert.NoError(t, err)
}

func deleteFile(t *testing.T, fileName string) {
	err := os.Remove(fileName)
	assert.NoError(t, err)
}
