package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpuriol/cuadrator/info"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Añade un participante a la actividad",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		participante := strings.Join(args, " ")

		err := info.AñadirParticipante(participante)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
