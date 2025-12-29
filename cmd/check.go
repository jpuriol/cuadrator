package cmd

import (
	"fmt"

	"github.com/jpuriol/cuadrator/app"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if information is valid",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := app.LoadAll()
		if err != nil {
			return err
		}

		err = d.Quadrant.ValidateNames(d.Participants)
		if err != nil {
			return err
		}

		err = d.Quadrant.ValidateShifts()
		if err != nil {
			return err
		}

		fmt.Println("Cuadrante cuadra!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
