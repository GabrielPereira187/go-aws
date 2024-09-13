package configuration

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func GetAwsConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if err != nil {
		return aws.Config{}, errors.New("error")
	}

	return cfg, nil
}
