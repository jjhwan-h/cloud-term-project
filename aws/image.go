package aws

import (
	"context"

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
