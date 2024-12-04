package internal

import (
	"TermProject/aws"
	"bufio"
	"fmt"
	"os"
)

type Cli struct {
	on bool //작업 완료 여부
	sc bufio.Scanner
}

func NewCli() (*Cli, error) {
	return &Cli{
		on: false,
		sc: *bufio.NewScanner(os.Stdin),
	}, nil
}

func (cli Cli) Start() error {
	cli.sc.Split(bufio.ScanWords)

	aws, err := aws.NewAws()
	if err != nil {
		return err
	}
	cli.processAnswer(aws)
	return nil
}

func (cli Cli) processAnswer(aws *aws.Aws) {
	cli.checkStatus()
	choice := cli.getPromptChoice()

	switch choice {
	case listInstance:
		err := aws.ListInstances()
		handleError(err)
	case availableZones:
		aws.AvailableZones()
	case startInstance:
		fmt.Println("Enter instance id: ")
		id := cli.scanString()
		err := aws.StartInstance(id)
		handleError(err)
	case availableRegions:
		aws.AvailableRegions()
	case stopInstance:
		fmt.Println("Enter instance id: ")
		id := cli.scanString()
		err := aws.StopInstance(id)
		handleError(err)
	case createInstance:
		fmt.Println("Enter ami id: ")
		id := cli.scanString()
		err := aws.CreateInstance(id)
		handleError(err)
	case rebootInstance:
		fmt.Println("Enter instance id: ")
		id := cli.scanString()
		err := aws.RebootInstance(id)
		handleError(err)
	case listImages:
		aws.ListImages()
	case quit:
		fmt.Println("quit")
		return
	default:
		fmt.Println("wrong number")
	}

	cli.processAnswer(aws)
}

func (cli Cli) checkStatus() {
	for {
		if !cli.on {
			return
		}
	}
}
func (cli Cli) getPromptChoice() option {
	printMenu()
	return cli.scanInt()
}
