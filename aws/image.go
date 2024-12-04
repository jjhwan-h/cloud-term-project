package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/viper"
)

func (aws Aws) ListImages() error {
	fmt.Println("list images")
	input := &ec2.DescribeImagesInput{
		Owners: []string{viper.GetString("AWS_OWNER_ID")},
		// Filters: []types.Filter{
		// 	{Name: a.String("name"),
		// 		Values: []string{"htcondor-worker"}},
		// },
	}

	images, err := aws.ec2.DescribeImages(context.TODO(), input)
	if err != nil {
		return err
	}
	for _, image := range images.Images {
		fmt.Printf("AMI ID: %s, Name: %s\n", *image.ImageId, *image.Name)
	}
	return nil
}
