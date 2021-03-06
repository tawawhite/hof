package app

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var pullLong = `Pull a Studios app.`

var PullCmd = &cobra.Command{

	Use: "pull <name or id>",

	Short: "Pull a Studios app.",

	Long: pullLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Ident'")
			cmd.Usage()
			os.Exit(1)
		}

		var ident string

		if 0 < len(args) {

			ident = args[0]

		}

		// Default body

		fmt.Println("hof studios app pull", ident)

	},
}
