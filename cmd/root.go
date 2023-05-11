// package
package cmd

import (
	"fmt"
	"os"

	"github.com/pipikai/yun/cmd/flags"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yun",
	Short: "a distribution file storage supprt multiple storage.",
	Long: `a distribution file storage supprt multiple storage,
		Complete documentation is available at xxx`,
}

// Execute app enter
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.ConfigFile, "f", "storage-1", "config file path")
}
