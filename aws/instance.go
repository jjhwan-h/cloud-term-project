package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (aws Aws) ListInstances() error {
	fmt.Println("Listing instances...")

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
func (aws Aws) CreateInstance(id string) error {
	return nil
}
func (aws Aws) RebootInstance() {
	fmt.Println("reboot instance")
}
func (aws Aws) StartInstance() {
	fmt.Println("start instance")
}
