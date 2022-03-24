package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpuriol/cuadrator/info"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the occupation a partipant has on a specif shift",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")

		occupations, err := info.GetPartipantOccupation(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, o := range occupations {
			fmt.Printf("%s: %s\n", o.ShiftName, o.OccupationName)
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
