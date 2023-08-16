package main

import (
	"fmt"
	"go-mem-db/tools/cli/cmd"
	"go-mem-db/tools/cli/db"
	"os"
)

func main() {
	db.Init("datastore")
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
