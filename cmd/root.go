package cmd

import (
	"fmt"
	"os"

	"github.com/crftd-tech/trunkver/internal"
	"github.com/crftd-tech/trunkver/internal/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: internal.Version,
	Use:     "trunkver [flags]",
	Short:   "trunkver generates versions for trunk-based apps",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.LogVerbose("trunkver version %s\n", internal.Version)
	},
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.PersistentFlags().BoolVarP(&log.Verbose, "verbose", "v", false, "Enable verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
