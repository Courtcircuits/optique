package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Bootstrap a new Optique module",
	Long:  `Generates a new Optique module`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}
