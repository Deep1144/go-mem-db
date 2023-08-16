/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-mem-db/tools/cli/db"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the key value pair from the database with the given key",
	Long: `Delete the key value pair from the database with the given key
Example: 

memdb delete -k key1`,
	Run: func(cmd *cobra.Command, args []string) {
		if db.DB == nil {
			fmt.Println("No database found. Please initialise a new database first")
			return
		}
		key, _ := cmd.Flags().GetString("key")
		db.DB.Delete(key)
		fmt.Printf("Deleted key %s\n", key)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("key", "k", "", "key to get the value of")

	deleteCmd.MarkFlagRequired("key")
}
