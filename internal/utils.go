package internal

import (
	"log"
	"strconv"
)

func (cli Cli) scanInt() option {
	cli.sc.Scan()
	v, _ := strconv.Atoi(cli.sc.Text())
	return option(v)
}

func (cli Cli) scanString() string {
	cli.sc.Scan()
	return cli.sc.Text()
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
