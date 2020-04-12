package config

import "github.com/kelseyhightower/envconfig"

type AWSConfig struct {
	Region          string `required:"true" envconfig:"region"`
	AccessKeyID     string `required:"true" envconfig:"access_key_id"`
	SecretAccessKey string `required:"true" envconfig:"secret_access_key"`
}

func MustLoad() AWSConfig {
	awsConfig := AWSConfig{}
	envconfig.MustProcess("AWS", &awsConfig)
	return awsConfig
}
