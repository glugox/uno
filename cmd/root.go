package cmd

import (
	"fmt"
	"os"

	uno "github.com/glugox/uno/pkg"
	"github.com/spf13/cobra"
)

var RootVerbose bool
var CMDLogger = uno.DefaultLogFactory().NewLogger()

var rootCmd = &cobra.Command{
	Use:   "uno",
	Short: "Uno is a web framework for easily creation of APIs",
	Long: `A Fast and with minimal dependencies web framework 
				  SQL database migrations
				  Complete documentation is available at http://uno.ba/docs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Uno! To run the server run: `go run main.go serve`, or set CMD_DEFAULT=serve in .env file")
	},
}

func Execute() {
	// Try to attach the default specified
	// command to main command args
	cDft := uno.Env("CMD_DEFAULT", "")
	if cDft != "" {
		rootCmd.SetArgs([]string{cDft})
	}

	// Execute main
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
