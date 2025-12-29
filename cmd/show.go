package cmd

import (
	"fmt"
	"strings"

	"github.com/jpuriol/cuadrator/app"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the occupation a participant has on a specif shift",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := app.LoadAll()
		if err != nil {
			return err
		}

		name := strings.Join(args, " ")

		if !d.Participants.Exists(name) {
			return fmt.Errorf("participant %q is not on participants file", name)
		}

		occupations := d.Quadrant.GetOccupation(name)

		if len(occupations) == 0 {
			fmt.Printf("No occupations for participant %q\n", name)
			return nil
		}

		for _, o := range occupations {
			fmt.Printf("%v: %q\n", o.ShiftID, d.Schema.OccupationName(o.OccupationID))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
