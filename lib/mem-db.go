package memdb

import (
	"encoding/json"
	"os"
	"sync"
)

type KeyValueDB struct {
	store  map[string]string
	dbName string
	mx     sync.RWMutex
}

func NewKeyValueDB(dbName string) *KeyValueDB {
	db := &KeyValueDB{
		store:  make(map[string]string),
		dbName: dbName,
	}
	db.loadData()

	return db
}

func (kvd *KeyValueDB) Put(key, value string) {
	kvd.mx.Lock()
	defer kvd.mx.Unlock()

	kvd.store[key] = value
	kvd.saveData()
}

func (kvd *KeyValueDB) Get(key string) (string, bool) {
	kvd.mx.RLock()
	defer kvd.mx.RUnlock()

	value, ok := kvd.store[key]
	return value, ok
}

func (kvd *KeyValueDB) Delete(key string) {
	kvd.mx.Lock()
	defer kvd.mx.Unlock()

	delete(kvd.store, key)
	kvd.saveData()
}

func (kvd *KeyValueDB) saveData() {
	data, _ := json.Marshal(kvd.store)
	os.WriteFile(kvd.dbName, data, os.ModePerm)
}

func (kvd *KeyValueDB) loadData() {
	if _, err := os.Stat(kvd.dbName); os.IsNotExist(err) {
		return
	}

	bytesData, _ := os.ReadFile(kvd.dbName)
	json.Unmarshal(bytesData, &kvd.store)
}
