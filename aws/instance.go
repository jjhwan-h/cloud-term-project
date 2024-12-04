package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
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
func (aws Aws) StopInstance(id string, dryRun bool) error {
	_, err := aws.ec2.StopInstances(context.TODO(), &ec2.StopInstancesInput{
		DryRun:      ToBool(dryRun),
		InstanceIds: []string{id},
	})
	if err != nil {
		if dryRun && strings.Contains(err.Error(), "DryRunOperation") {
			return nil
		}
		return err
	}
	return nil
}
func (aws Aws) CreateInstance(id string) error {
	fmt.Println("Creating...")
	res, err := aws.ec2.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		MaxCount:     ToInt32(1),
		MinCount:     ToInt32(1),
		ImageId:      ToString(id),
		InstanceType: types.InstanceTypeT2Micro,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Successfully started EC2 instance %s based on AMI %s\n",
		*res.ReservationId, id)
	return nil
}
func (aws Aws) RebootInstance() {
	fmt.Println("reboot instance")
}
func (aws Aws) StartInstance(id string, dryRun bool) error {
	_, err := aws.ec2.StartInstances(context.TODO(), &ec2.StartInstancesInput{
		DryRun:      ToBool(dryRun),
		InstanceIds: []string{id},
	})
	if err != nil {
		if dryRun && strings.Contains(err.Error(), "DryRunOperation") {
			return nil
		}
		return err
	}
	return nil
}
