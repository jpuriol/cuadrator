/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jpuriol/cuadrator/data"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show how many participants we have per shift",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := data.LoadAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, shiftID := range d.Quadrant.OrderedShiftIDs() {
			fmt.Printf("%d -> [%s]\n", shiftID, d.Schema.ShiftName(shiftID))
			shift := d.Quadrant[shiftID]
			for _, occupationID := range shift.OrderedOccupationIDs() {
				teams := shift[occupationID]
				fmt.Printf(" %s: %d\n", d.Schema.OccupationName(occupationID), len(teams))
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
