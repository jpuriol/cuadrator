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
        quadrant, err := data.ReadQuadrant()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

        shifts, err := data.ReadSchema()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

		for  _, shiftID := range quadrant.OrderedShiftIDs() {
            fmt.Printf("%d -> [%s]\n", shiftID, shifts.ShiftName(shiftID))
            shift := quadrant[shiftID]
            for _, ocuppationID := range shift.OrderedOcuppationIDs() {
                teams := shift[ocuppationID]
                fmt.Printf(" %d: %s (%d)\n", len(teams), shifts.OcupationName(ocuppationID), ocuppationID)
            }

        }
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
