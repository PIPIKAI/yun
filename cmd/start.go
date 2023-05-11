package cmd

import (
	"github.com/pipikai/yun/core"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
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
	rootCmd.AddCommand(startCmd)
}
