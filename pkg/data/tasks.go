package data

import (
	"errors"
	"fmt"
	"sync"
)

var (
	db   TasksDB
	once sync.Once
)

type TaskID string

type Task struct {
	ID          TaskID `json:"id"`
	Done        bool   `json:"done"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// CreateOn    time.Time
}

type TasksDB struct {
	mu     sync.Mutex
	bucket map[string]Task
}

func NewTasksDB() *TasksDB {
	once.Do(func() {
		db = TasksDB{
			bucket: make(map[string]Task),
		}
	})
	return &db
}

func (db *TasksDB) Create(s Task) error {
	if _, ok := db.bucket[string(s.ID)]; ok {
		return errors.New("user already exist")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.bucket[string(s.ID)] = s
	return nil
}

func (db *TasksDB) Read(id TaskID) (Task, error) {
	if c, ok := db.bucket[string(id)]; ok {
		return c, nil
	}

	return Task{}, fmt.Errorf("entity with %s id, doesnt exist", id)
}

func (db *TasksDB) Update(s Task) error {
	_, err := db.Read(s.ID)
	if err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.bucket[string(s.ID)] = s
	return nil
}

func (db *TasksDB) Delete(id TaskID) error {
	_, err := db.Read(id)
	if err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.bucket, string(id))
	return nil
}

func (db *TasksDB) List() []Task {
	x := []Task{}

	for _, v := range db.bucket {
		x = append(x, v)
	}
	return x
}
