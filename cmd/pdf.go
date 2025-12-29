package cmd

import (
	"fmt"
	"os"

	"github.com/jpuriol/cuadrator/exporter"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Generate a pdf document with the cuadrant information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := exporter.PrintPDF()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("PDF created!")

	},
}

func init() {
	rootCmd.AddCommand(pdfCmd)
}
