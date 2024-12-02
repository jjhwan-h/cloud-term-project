package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/viper"
)

type Aws struct {
	cfg config.Config
	ec2 ec2.Client
}

func NewAws() (*Aws, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			viper.GetString("AWS_ACCESS_KEY_ID"),
			viper.GetString("AWS_SECRET_ACCESS_KEY"),
			"",
		)),
		config.WithRegion(viper.GetString("AWS_REGION")))
	if err != nil {
		return nil, err
	}
	return &Aws{
		cfg: cfg,
		ec2: *ec2.NewFromConfig(cfg),
	}, nil
}
