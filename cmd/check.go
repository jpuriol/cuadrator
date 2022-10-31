package cmd

import (
	"fmt"
	"os"

	"github.com/jpuriol/cuadrator/data"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if information is valid",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		quadrant, err := data.ReadQuadrant()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = quadrant.Check()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Cuadrante cuadra!")

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
