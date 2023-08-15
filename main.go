package main

import (
	"go-db/lib"
)

func main() {
	db := lib.NewKeyValueDB("datastore.json")
	db.Put("try2", "1")

	db.Get("try")
	db.Delete("try1")
}
