package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func Config() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithDefaultRegion("us-east-1"),
	)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}
