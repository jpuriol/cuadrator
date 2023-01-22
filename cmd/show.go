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

        participants, err := data.ReadParticipants()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

        shifts, err := data.ReadSchema()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

		name := strings.Join(args, " ")

        if !participants.Exists(name) {
            fmt.Fprintln(os.Stderr, "Participants is not on participants file")
            os.Exit(1)
        }

		occupations := quadrant.GetOcupation(name)

		for _, o := range occupations {
			fmt.Printf("%v: %q\n", o.ShifID, shifts.OcupationName(o.OccupationID))
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
