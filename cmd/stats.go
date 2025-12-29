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

		maxOccupationLen := 0
		for _, name := range d.Schema.Occupations {
			if len(name) > maxOccupationLen {
				maxOccupationLen = len(name)
			}
		}

		for _, shiftID := range d.Quadrant.OrderedShiftIDs() {
			shiftName := d.Schema.ShiftName(shiftID)
			header := fmt.Sprintf(" %s ", shiftName)
			separator := ""
			for i := 0; i < len(header); i++ {
				separator += "-"
			}
			fmt.Printf("%s\n", separator)
			fmt.Printf("%s\n", header)
			fmt.Printf("%s\n", separator)
			shift := d.Quadrant[shiftID]
			occupationIDs := shift.OrderedOccupationIDs()
			for i := 0; i < len(occupationIDs); i += 2 {
				line := ""
				for j := 0; j < 2 && i+j < len(occupationIDs); j++ {
					occupationID := occupationIDs[i+j]
					teams := shift[occupationID]
					count := len(teams)
					bar := ""
					for k := 0; k < count; k++ {
						bar += "█"
					}
					line += fmt.Sprintf(" %*s | %-12s (%d)  ", maxOccupationLen, d.Schema.OccupationName(occupationID), bar, count)
				}
				fmt.Println(line)
			}
			fmt.Println()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
