package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "memdb",
	Short: "In memory database in GO",
}
