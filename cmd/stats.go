/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jpuriol/cuadrator/app"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show how many participants we have per shift",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := app.LoadAll()
		if err != nil {
			return err
		}

		for _, shiftID := range d.Quadrant.OrderedShiftIDs() {
			fmt.Printf("%d -> [%s]\n", shiftID, d.Schema.ShiftName(shiftID))
			shift := d.Quadrant[shiftID]
			for _, occupationID := range shift.OrderedOccupationIDs() {
				teams := shift[occupationID]
				fmt.Printf(" %s: %d\n", d.Schema.OccupationName(occupationID), len(teams))
			}

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
