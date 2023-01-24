package database

import (
	"fmt"
	"io/ioutil"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

const path = "/opt/mem.db.json"

type MemDB struct {
	subsribers map[int64]interface{}
	mu         *sync.RWMutex
}

func NewMemDB() Database {
	db := &MemDB{
		subsribers: make(map[int64]interface{}),
		mu:         &sync.RWMutex{},
	}

	data, err := ioutil.ReadFile(path)
	if err == nil {
		if err := jsoniter.Unmarshal(data, &db.subsribers); err != nil {
			logrus.Panicf("Invalid unmarshal db: %s", err)
		}
	}

	return db
}

func (db *MemDB) GetSubscribers() ([]int64, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	slice := make([]int64, 0, len(db.subsribers))
	for subsriber := range db.subsribers {
		slice = append(slice, subsriber)
	}

	return slice, nil
}

func (db *MemDB) AddSubscriber(id int64) (bool, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, ok := db.subsribers[id]
	if !ok {
		db.subsribers[id] = nil
	}

	data, err := jsoniter.Marshal(db.subsribers)
	if err != nil {
		return false, fmt.Errorf("Invalid marshal db: %s", err)
	}

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return false, fmt.Errorf("Invalid write db: %s", err)
	}

	return ok, nil
}
