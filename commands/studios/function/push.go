package function

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var pushLong = `Push a Studios function.`

var PushCmd = &cobra.Command{

	Use: "push <name or id>",

	Short: "Push a Studios function.",

	Long: pushLong,

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

		fmt.Println("hof studios function push", ident)

	},
}
