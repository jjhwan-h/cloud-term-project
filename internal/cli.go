package internal

import (
	"TermProject/aws"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		if err != nil {
			log.Println(err)
		}
	case availableZones:
		aws.AvailableZones()
	case startInstance:
		aws.StartInstance()
	case availableRegions:
		aws.AvailableRegions()
	case stopInstance:
		aws.StopInstance()
	case createInstance:
		id := cli.scanString()
		err := aws.CreateInstance(id)
		if err != nil {
			log.Println(err)
		}
	case rebootInstance:
		aws.RebootInstance()
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

func (cli Cli) scanInt() option {
	cli.sc.Scan()
	v, _ := strconv.Atoi(cli.sc.Text())
	return option(v)
}

func (cli Cli) scanString() string {
	cli.sc.Scan()
	return cli.sc.Text()
}
