package cmd

import (
	"fmt"
	"os"

	"github.com/crftd-tech/trunkver/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: internal.Version,
	Use:     "trunkver [flags]",
	Short:   "trunkver generates versions for trunk-based apps",
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
