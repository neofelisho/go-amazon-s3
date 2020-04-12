package helper

import "github.com/aws/aws-sdk-go/service/s3"

func GetNamesOfBucket(buckets []*s3.Bucket) []string {
	var bs []string
	for _, b := range buckets {
		bs = append(bs, *b.Name)
	}
	return bs
}
