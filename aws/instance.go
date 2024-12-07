package aws

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
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
					addr := ""
					name := ""
					if instance.PublicIpAddress != nil {
						addr = *instance.PublicIpAddress
					}
					for _, tag := range instance.Tags {
						if *tag.Key == "Name" {
							name = *tag.Value
							break
						}
					}
					rows = append(rows, table.Row{*instance.InstanceId,
						*instance.ImageId,
						string(instance.InstanceType),
						string(instance.State.Name),
						string(instance.Monitoring.State),
						name,
						addr})

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
func (aws Aws) StopInstance(ch []string) (*string, error) {
	//DryRun : 요청 유효성 및 잠재적인 오류 확인
	err := aws.stopInstance(ch[0], true)
	if err != nil {
		return nil, err
	}
	//Run
	err = aws.stopInstance(ch[0], false)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully stopped instance %s\n", ch[0])

	return ptr(res), nil
}
func (aws Aws) CreateInstance(ch []string) (*string, error) {
	keyFilePath := viper.GetString("PRIVATE_KEY_PATH")
	fileName := filepath.Base(keyFilePath)
	keyName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	instanceOutput, err := aws.ec2.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		MaxCount:         ToInt32(1),
		MinCount:         ToInt32(1),
		ImageId:          ToString(ch[0]),
		InstanceType:     types.InstanceTypeT2Micro,
		SecurityGroupIds: []string{ch[1]},
		KeyName:          &keyName,
	})
	if err != nil {
		return nil, err
	}

	res := fmt.Sprintf("Successfully started EC2 instance %s based on AMI %s\n",
		*instanceOutput.ReservationId, ch[0])
	return ptr(res), nil
}
func (aws Aws) RebootInstance(ch []string) (*string, error) {
	// DryRun : 요청 유효성 및 잠재적인 오류 확인
	err := aws.rebootInstance(ch[0], true)
	if err != nil {
		return nil, err
	}
	// Run
	err = aws.rebootInstance(ch[0], false)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully rebooted instance %s\n", ch[0])
	return ptr(res), nil
}
func (aws Aws) StartInstance(ch []string) (*string, error) {
	//DryRun : 요청 유효성 및 잠재적인 오류 확인
	err := aws.startInstance(ch[0], true)
	if err != nil {
		return nil, err
	}
	//Run
	err = aws.startInstance(ch[0], false)
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully started instance %s\n", ch[0])

	return ptr(res), nil
}
func (aws Aws) ConnectInstance(ch []string) (*ssh.Client, error) {
	privateKeyPath := viper.GetString("PRIVATE_KEY_PATH")
	user := viper.GetString("USER")
	key, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         15 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", ch[0]+":22", config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (aws Aws) TerminateInstance(ch []string) (*string, error) {
	if ch[1] == "main" {
		err := errors.New("main instance cannot be terminated")
		return nil, err
	}
	_, err := aws.ec2.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{ch[0]},
	})
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully terminated EC2 instance %s\n", ch[0])
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
