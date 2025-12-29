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
		d, err := data.LoadAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		name := strings.Join(args, " ")

		if !d.Participants.Exists(name) {
			fmt.Fprintf(os.Stderr, "Participant %q is not on participants file\n", name)
			os.Exit(1)
		}

		occupations := d.Quadrant.GetOccupation(name)

		if len(occupations) == 0 {
			fmt.Printf("No occupations for participant %q\n", name)
			return
		}

		for _, o := range occupations {
			fmt.Printf("%v: %q\n", o.ShiftID, d.Schema.OccupationName(o.OccupationID))
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
