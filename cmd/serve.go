package cmd

import (
	uno "github.com/glugox/uno/pkg"
	"github.com/spf13/cobra"
)

var verbose bool

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the Uno server",
	Run: func(cmd *cobra.Command, args []string) {
		uno, err := uno.New()
		if err != nil {
			CMDLogger.Error(err.Error())
			return
		}

		// Initialize DB, migrations, etc
		err = uno.Init()
		if err != nil {
			CMDLogger.Error(err.Error())
			return
		}

		uno.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
