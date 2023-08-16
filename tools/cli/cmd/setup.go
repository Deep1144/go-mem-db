package cmd

import (
	"fmt"
	"go-mem-db/tools/cli/db"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up the database",
	Long: `This command initializes the database and creates the tables required for the application to run.
Example:
memdb setup --database-name datastore`,
	Run: func(cmd *cobra.Command, args []string) {
		databaseName, _ := cmd.Flags().GetString("database-name")
		db.Init(databaseName)

		fmt.Println("Database setup completed")
	},
}

func init() {
	RootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringP("database-name", "n", "", "Database name")

	setupCmd.MarkFlagRequired("database-name")
}
