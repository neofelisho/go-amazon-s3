package helper

import "regexp"

const patternBucketName string = `^(([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])\.)*([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])$`

func ValidateBucketName(name string) bool {
	if len(name) < 3 || len(name) > 63 {
		return false
	}
	matched, err := regexp.MatchString(patternBucketName, name)
	if err != nil {
		return false
	}
	return matched
}
