package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/viper"
)

func (aws Aws) ListImages() ([]table.Row, error) {
	input := &ec2.DescribeImagesInput{
		Owners: []string{viper.GetString("AWS_OWNER_ID")},
		// Filters: []types.Filter{
		// 	{Name: a.String("name"),
		// 		Values: []string{"htcondor-worker"}},
		// },
	}

	images, err := aws.ec2.DescribeImages(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	var rows []table.Row
	for _, image := range images.Images {
		rows = append(rows, table.Row{
			*image.ImageId, *image.Name,
		})
	}
	return rows, nil
}

func (aws Aws) CreateImage(ch []string) (*string, error) {
	createOutput, err := aws.ec2.CreateImage(context.TODO(), &ec2.CreateImageInput{
		InstanceId: ToString(ch[len(ch)-2]),
		Name:       ToString(ch[len(ch)-1]),
	})
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully created Image\n Id: %s, Name: %s\n", *createOutput.ImageId, ch[len(ch)-1])
	return ptr(res), nil
}

func (aws Aws) DeleteImage(ch []string) (*string, error) {
	_, err := aws.ec2.DeregisterImage(context.TODO(), &ec2.DeregisterImageInput{
		ImageId: ToString(ch[0]),
	})
	if err != nil {
		return nil, err
	}
	res := fmt.Sprintf("Successfully deleted Image\n Id: %s\n", ch[0])
	return ptr(res), nil
}
