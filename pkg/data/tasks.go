package data

import (
	"errors"
	"fmt"
	"sync"
)

type Store interface {
	Create(*Task) error
	Read(TaskID) (*Task, error)
	Update(*Task) error
	Delete(TaskID) error
	List() []*Task
}

type TaskID string

// Task represents simple TODO task entity
//
// swagger:model
type Task struct {
	// the id of the task
	//
	// Required: false
	// min: 1
	ID TaskID `json:"id"`

	// the status of the task
	//
	// required: false
	// example: false
	Done bool `json:"done"`

	// the title of the task
	// required: true
	// max length: 50
	Title string `json:"title"`

	// tasks short description
	// required: true
	// max length: 255
	Description string `json:"description"`
	// CreateOn    time.Time
}

// TaskDB acts as a database
type TasksDB struct {
	mu     sync.Mutex
	bucket map[string]*Task
}

func NewTasksDB() *TasksDB {
	return &TasksDB{
		bucket: make(map[string]*Task),
		mu:     sync.Mutex{},
	}
}

func (db *TasksDB) Create(s *Task) error {
	if _, ok := db.bucket[string(s.ID)]; ok {
		return errors.New("user already exist")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.bucket[string(s.ID)] = s
	return nil
}

func (db *TasksDB) Read(id TaskID) (*Task, error) {
	if c, ok := db.bucket[string(id)]; ok {
		return c, nil
	}

	return nil, fmt.Errorf("entity with %s id, doesnt exist", id)
}

func (db *TasksDB) Update(s *Task) error {
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

func (db *TasksDB) List() []*Task {
	x := []*Task{}

	for _, v := range db.bucket {
		x = append(x, v)
	}
	return x
}
