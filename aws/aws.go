package aws

import "fmt"

type Aws struct {
}

func NewAws() (*Aws, error) {
	return &Aws{}, nil
}

func (aws Aws) ListInstance() {
	fmt.Println("list instance")
	return
}
func (aws Aws) StopInstance() {
	fmt.Println("stop instance")
	return
}
func (aws Aws) CreateInstance() {
	fmt.Println("create instance")
	return
}
func (aws Aws) RebootInstance() {
	fmt.Println("reboot instance")
	return
}
func (aws Aws) StartInstance() {
	fmt.Println("start instance")
	return
}
func (aws Aws) AvailableZones() {
	fmt.Println("available zones")
	return
}
func (aws Aws) AvailableRegions() {
	fmt.Println("available regions")
	return
}
func (aws Aws) ListImages() {
	fmt.Println("list images")
	return
}
