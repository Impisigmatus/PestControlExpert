package database

import "sync"

type MemDB struct {
	subsribers map[int64]interface{}
	mu         *sync.RWMutex
}

func NewMemDB() Database {
	return &MemDB{
		subsribers: make(map[int64]interface{}),
		mu:         &sync.RWMutex{},
	}
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

func (db *MemDB) AddSubscriber(id int64) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.subsribers[id] = nil
	return nil
}
