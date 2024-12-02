package aws

import (
	"context"
	"fmt"

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
	// options := ec2.Options{Region: viper.GetString("AWS_REGION")}
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

func (aws Aws) ListInstance() error {
	fmt.Println("list instance")

	input := &ec2.DescribeInstancesInput{}
	for {
		instances, err := aws.ec2.DescribeInstances(context.TODO(), input)
		if err != nil {
			return err
		}

		for _, r := range instances.Reservations {
			for _, instance := range r.Instances {
				fmt.Printf(
					"[id] %s, "+
						"[AMI] %s, "+
						"[type] %s, "+
						"[state] %10s, "+
						"[monitoring state] %s\n",
					*instance.InstanceId,
					*instance.ImageId,
					instance.InstanceType,
					instance.State.Name,
					instance.Monitoring.State,
				)
			}
		}
		if instances.NextToken == nil {
			break
		}
		input.NextToken = instances.NextToken
	}

	return nil
}
func (aws Aws) StopInstance() {
	fmt.Println("stop instance")
}
func (aws Aws) CreateInstance() {
	fmt.Println("create instance")
}
func (aws Aws) RebootInstance() {
	fmt.Println("reboot instance")
}
func (aws Aws) StartInstance() {
	fmt.Println("start instance")
}
func (aws Aws) AvailableZones() {
	fmt.Println("available zones")
}
func (aws Aws) AvailableRegions() {
	fmt.Println("available regions")
}
func (aws Aws) ListImages() {
	fmt.Println("list images")
}
