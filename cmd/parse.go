package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/crftd-tech/trunkver/internal/trunkver"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"p"},
	Short:   "Parse a TrunkVer",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var inputString string

		if len(args) > 0 {
			var err error
			inputString = args[0]
			if err != nil {
				panic(err)
			}
		} else {
			reader := bufio.NewReader(cmd.InOrStdin())
			inputBA, _, err := reader.ReadLine()
			inputString = string(inputBA)
			if err != nil {
				panic(err)
			}
		}

		version, err := trunkver.ParseTrunkVer(inputString)
		if err != nil {
			panic(err)
		}

		if format, _ := cmd.Flags().GetString("format"); format != "" {
			tpl := template.Must(template.New("trunkver").Parse(format))
			err := tpl.Execute(cmd.OutOrStdout(), version)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(cmd.OutOrStdout())
		} else {
			jsonStr, err := json.Marshal(version)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(cmd.OutOrStdout(), string(jsonStr))
		}
	},
}

func init() {
	parseCmd.Flags().StringP("format", "f", "", "go template format string to use for output")
	rootCmd.AddCommand(parseCmd)
}
