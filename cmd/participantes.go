package cmd

import (
	"fmt"
	"os"

	"github.com/jpuriol/cuadrator/info"
	"github.com/spf13/cobra"
)

var participantes = &cobra.Command{
	Use:   "participantes",
	Short: "Muestra los participantes inscritos en la actividad",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		particpantes, err := info.VerParticipantes()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, p := range particpantes {
			fmt.Println("  ", p)
		}
	},
}

func init() {
	rootCmd.AddCommand(participantes)
}
