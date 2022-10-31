package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpuriol/cuadrator/data"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the occupation a partipant has on a specif shift",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quadrant, err := data.ReadQuadrant()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		name := strings.Join(args, " ")
		occupations, err := quadrant.GetOcupation(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, o := range occupations {
			fmt.Printf("%v: %q\n", o.ShifID, o.OccupationName)
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
