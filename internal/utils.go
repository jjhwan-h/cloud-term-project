package internal

import (
	"log"
	"os"
	"os/exec"
)

func handleResult(res *string, err error) *string {
	if err != nil {
		log.Println(err)
		return ptr(err.Error())
	}
	if res != nil {
		return res
	}
	return nil
}

func ptr(s string) *string {
	return &s
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
