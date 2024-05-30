package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/crftd-tech/trunkver/internal"
	"github.com/crftd-tech/trunkver/internal/trunkver"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:     "parse [flags] [trunkver | -]",
	Long:    `Parses a TrunkVer and outputs it as JSON. If a format string is provided, it will be used instead.`,
	Aliases: []string{"p"},
	Short:   "Parse a TrunkVer",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var inputString string

		if args[0] == "-" {
			reader := bufio.NewReader(cmd.InOrStdin())
			inputBA, _ := internal.Must2(reader.ReadLine())
			inputString = string(inputBA)
		} else {
			inputString = args[0]
		}

		version := internal.Must(trunkver.ParseTrunkVer(inputString))

		if format, _ := cmd.Flags().GetString("format"); format != "" {
			tpl := template.Must(template.New("trunkver").Parse(format))
			internal.Must(tpl.Execute(cmd.OutOrStdout(), version), nil)
			fmt.Fprintln(cmd.OutOrStdout())
		} else {
			jsonStr := internal.Must(json.Marshal(version))
			fmt.Fprintln(cmd.OutOrStdout(), string(jsonStr))
		}
	},
}

func init() {
	parseCmd.Flags().StringP("format", "f", "", "go template format string to use for output")
	rootCmd.AddCommand(parseCmd)
}
