package cmd

import (
	"github.com/pipikai/yun/core"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start processes",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	core.Start()
}

func init() {
	RootCmd.AddCommand(StartCmd)
}
