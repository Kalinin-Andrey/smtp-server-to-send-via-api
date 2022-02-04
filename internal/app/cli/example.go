package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "This is example command.",
	Long:  `This is the long description for the example command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("example called")
	},
}

func init() {
	app.rootCmd.AddCommand(exampleCmd)
}
