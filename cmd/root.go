package cmd

import (
	"fmt"
	"os"

	"github.com/pipikai/yun/cmd/flags"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "yun",
	Short: "a distribution file storage supprt multiple storage.",
	Long: `a distribution file storage supprt multiple storage,
		Complete documentation is available at xxx`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&flags.ConfigFile, "f", "storage-1", "config file path")
}
