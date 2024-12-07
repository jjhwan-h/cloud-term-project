package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/charmbracelet/bubbles/table"
)

func (aws Aws) AvailableZones() ([]table.Row, error) {
	zones, err := aws.ec2.DescribeAvailabilityZones(context.TODO(), &ec2.DescribeAvailabilityZonesInput{})

	if err != nil {
		return nil, err
	}

	var rows []table.Row
	for _, zone := range zones.AvailabilityZones {
		rows = append(rows, table.Row{*zone.ZoneId, *zone.RegionName, *zone.ZoneName})
	}

	return rows, nil
}
func (aws Aws) AvailableRegions() ([]table.Row, error) {
	regions, err := aws.ec2.DescribeRegions(context.TODO(), &ec2.DescribeRegionsInput{})

	if err != nil {
		return nil, err
	}

	var rows []table.Row
	for _, region := range regions.Regions {
		rows = append(rows, table.Row{*region.RegionName, *region.Endpoint})
	}

	return rows, nil
}

func (aws Aws) ListSecurityGroup() ([]table.Row, error) {
	sgs, err := aws.ec2.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{})

	if err != nil {
		return nil, err
	}

	var rows []table.Row
	for _, sg := range sgs.SecurityGroups {
		rows = append(rows, table.Row{*sg.GroupId, *sg.GroupName})
	}
	return rows, nil
}
