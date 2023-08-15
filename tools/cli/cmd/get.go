package cmd

import (
	"fmt"
	"go-mem-db/tools/cli/db"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of the key",
	Long: `Simple get command to get the value of the provided key. 
Example:

memdb get --key johnny`,
	Run: func(cmd *cobra.Command, args []string) {
		if db.DB == nil {
			fmt.Println("No database found. Please initialise a new database first")
			return
		}
		key, _ := cmd.Flags().GetString("key")
		value, ok := db.DB.Get(key)
		if !ok {
			fmt.Printf("No value found for key %s\n", key)
			return
		}

		fmt.Println(value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("key", "k", "", "key to get the value of")

	getCmd.MarkFlagRequired("key")
}
