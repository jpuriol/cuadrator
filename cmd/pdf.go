package cmd

import (
	"fmt"

	"github.com/jpuriol/cuadrator/app"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Generate a pdf document with the cuadrant information",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := app.GeneratePDF()
		if err != nil {
			return err
		}

		fmt.Println("PDF created!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pdfCmd)
}
