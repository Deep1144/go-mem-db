package main

import (
	"flag"
	"fmt"
	memdb "go-mem-db/lib"
)

var (
	action string
	key    string
	value  string
)

func init() {
	flag.StringVar(&action, "a", "", "Action to perform (get, put, delete)")
	flag.StringVar(&key, "k", "", "Key for the action")
	flag.StringVar(&value, "v", "", "Value to put for the action")
	flag.Parse()
}

func main() {
	db := memdb.NewKeyValueDB("datastore")

	if key == "" {
		fmt.Printf("Invalid Key.")
		return
	}

	switch action {
	case "get":
		val, ok := db.Get(key)
		if ok {
			fmt.Println(val)
		}
	case "put":
		db.Put(key, value)
		fmt.Printf("Inserted the key %s with value %s into the DB. ", key, value)
	case "delete":
		db.Delete(key)
		fmt.Printf("Deleted the key %s from the DB. ", key)
	default:
		fmt.Println("Invalid action. supported actions are get, put, delete.")
	}
}
