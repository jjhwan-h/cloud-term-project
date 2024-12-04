package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (aws Aws) AvailableZones() error {
	fmt.Println("Available zones...")
	zones, err := aws.ec2.DescribeAvailabilityZones(context.TODO(), &ec2.DescribeAvailabilityZonesInput{})

	if err != nil {
		return err
	}

	for _, zone := range zones.AvailabilityZones {
		fmt.Printf(
			"[id] %s,  "+
				"[region] %15s,  "+
				"[zone] %15s\n", *zone.ZoneId, *zone.RegionName, *zone.ZoneName)
	}

	return nil
}
func (aws Aws) AvailableRegions() error {
	fmt.Println("Available regions...")
	regions, err := aws.ec2.DescribeRegions(context.TODO(), &ec2.DescribeRegionsInput{})

	if err != nil {
		return err
	}

	for _, region := range regions.Regions {
		fmt.Printf(
			"[region] %15s, "+
				"[endpoint] %s\n",
			*region.RegionName,
			*region.Endpoint)
	}

	return nil
}
