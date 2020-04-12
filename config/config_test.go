package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustLoad(t *testing.T) {
	awsConfig := MustLoad()
	assert.Greater(t, awsConfig.Region, "")
	assert.Greater(t, awsConfig.AccessKeyID, "")
	assert.Greater(t, awsConfig.SecretAccessKey, "")
}
