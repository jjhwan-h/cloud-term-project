package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/charmbracelet/bubbles/table"
)

type (
	state   *string
	RUNNING state
	STOPPED state
)

const (
	DRYOPERATION = "DryRunOperation"
)

func (aws Aws) ListInstances(state *string) ([]table.Row, error) {
	var rows []table.Row
	input := &ec2.DescribeInstancesInput{}
	for {
		instances, err := aws.ec2.DescribeInstances(context.TODO(), input)
		if err != nil {
			return nil, err
		}

		for _, r := range instances.Reservations {
			for _, instance := range r.Instances {
				if state == nil || *state == string(instance.State.Name) {
					rows = append(rows, table.Row{*instance.InstanceId,
						*instance.ImageId,
						string(instance.InstanceType),
						string(instance.State.Name),
						string(instance.Monitoring.State)})
				}
			}
		}
		if instances.NextToken == nil {
			break
		}
		input.NextToken = instances.NextToken
	}

	return rows, nil
}
func (aws Aws) StopInstance(id string) (*string, error) {
	//DryRun : 요청 유효성 및 잠재적인 오류 확인
	err := aws.stopInstance(id, true)
	if err != nil {
		return nil, err
	}
	//Run
	err = aws.stopInstance(id, false)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully stop instance %s\n", id)

	return ptr(res), nil
}
func (aws Aws) CreateInstance(id string) (*string, error) {
	instanceOutput, err := aws.ec2.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		MaxCount:     ToInt32(1),
		MinCount:     ToInt32(1),
		ImageId:      ToString(id),
		InstanceType: types.InstanceTypeT2Micro,
	})
	if err != nil {
		return nil, err
	}

	res := fmt.Sprintf("Successfully started EC2 instance %s based on AMI %s\n",
		*instanceOutput.ReservationId, id)
	return ptr(res), nil
}
func (aws Aws) RebootInstance(id string) (*string, error) {
	// DryRun : 요청 유효성 및 잠재적인 오류 확인
	err := aws.rebootInstance(id, true)
	if err != nil {
		return nil, err
	}
	// Run
	err = aws.rebootInstance(id, false)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully reboot instance %s\n", id)
	return ptr(res), nil
}
func (aws Aws) StartInstance(id string) (*string, error) {
	//DryRun : 요청 유효성 및 잠재적인 오류 확인
	err := aws.startInstance(id, true)
	if err != nil {
		return nil, err
	}
	//Run
	err = aws.startInstance(id, false)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully start instance %s\n", id)

	return ptr(res), nil
}
func (aws Aws) startInstance(id string, dryRun bool) error {
	_, err := aws.ec2.StartInstances(context.TODO(), &ec2.StartInstancesInput{
		DryRun:      ToBool(dryRun),
		InstanceIds: []string{id},
	})
	if err != nil {
		if dryRun && strings.Contains(err.Error(), DRYOPERATION) {
			return nil
		}
		return err
	}
	return nil
}
func (aws Aws) stopInstance(id string, dryRun bool) error {
	_, err := aws.ec2.StopInstances(context.TODO(), &ec2.StopInstancesInput{
		DryRun:      ToBool(dryRun),
		InstanceIds: []string{id},
	})
	if err != nil {
		if dryRun && strings.Contains(err.Error(), DRYOPERATION) {
			return nil
		}
		return err
	}
	return nil
}
func (aws Aws) rebootInstance(id string, dryRun bool) error {
	_, err := aws.ec2.RebootInstances(context.TODO(), &ec2.RebootInstancesInput{
		DryRun:      ToBool(dryRun),
		InstanceIds: []string{id},
	})
	if err != nil {
		if dryRun && strings.Contains(err.Error(), DRYOPERATION) {
			return nil
		}
		return err
	}
	return nil
}
