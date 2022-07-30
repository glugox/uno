package cmd

import (
	"github.com/spf13/cobra"
)

var (
	IsCreate bool
)

// serveCmd represents the serve command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs available DB migrations",
	Run: func(cmd *cobra.Command, args []string) {

		// // We can create modified config for our test app
		// cfg, err := uno.NewConfig()
		// if err != nil {
		// 	panic(err)
		// }

		// // Apply verbose from cli into app config
		// v := cmd.Flags().Lookup("verbose").Value.String()
		// if v == "true" {
		// 	cfg.App.Verbose = true
		// }

		// uno, err := uno.WithConfig(cfg)
		// if err != nil {
		// 	panic(err)
		// }

		// // Flags:
		// if IsCreate {
		// 	if len(args) == 0 {
		// 		panic("Please provide migration name as the 1st argument")
		// 	}

		// 	_, err := uno.DB.Migrator.CreateMigration(args[0])
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	// We only want to create migration, not to proceed with running, etc.
		// 	return
		// }

		// // Init wil do the migrations
		// err = uno.DB.Init()
		// if err != nil {
		// 	panic(err)
		// }
	},
}

func init() {
	migrateCmd.PersistentFlags().BoolVarP(&IsCreate, "create", "c", false, "Do we want to create migration file")
	migrateCmd.PersistentFlags().BoolVarP(&RootVerbose, "verbose", "v", false, "Verbose output")
	rootCmd.AddCommand(migrateCmd)
}
