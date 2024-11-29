package cmd

import (
	"TermProject/internal"
	"log"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Start a dynamic AWS management program in a CLI environment",
	Long:  "Start a dynamic AWS management program in a CLI environment",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := internal.NewCli()
		if err != nil {
			log.Fatal(err)
		}
		err = cli.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(cliCmd)
}
