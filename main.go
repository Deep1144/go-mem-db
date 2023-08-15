package main

import (
	memdb "go-mem-db/lib"
)

func main() {
	db := memdb.NewKeyValueDB("datastore")
	db.Put("try2", "1")

	db.Get("try")
	db.Delete("try1")
}
