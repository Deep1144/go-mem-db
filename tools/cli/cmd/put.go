package cmd

import (
	"fmt"
	"go-mem-db/tools/cli/db"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Save a key value pair",
	Long: `This command saves a key value pair in the database.
Example:

membd put --key "name" --value "John Doe"`,
	Run: func(cmd *cobra.Command, args []string) {
		if db.DB == nil {
			fmt.Println("No database found. Please initialise a new database first")
			return
		}
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")

		db.DB.Put(key, value)
		fmt.Printf("Successfully added key %s with value %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(putCmd)

	putCmd.Flags().StringP("key", "k", "", "key to store")
	putCmd.Flags().StringP("value", "v", "", "value to store for the key")

	putCmd.MarkFlagRequired("key")
	putCmd.MarkFlagRequired("value")
}
