package aws

import "fmt"

func (aws Aws) AvailableZones() {
	fmt.Println("available zones")
}
func (aws Aws) AvailableRegions() {
	fmt.Println("available regions")
}
