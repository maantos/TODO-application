package db

import (
	"errors"
	"sync"

	"github.com/maantos/todoApplication/pkg/domain"
)

var ErrTaskNotFound = errors.New("element not found")
var ErrTaskAlreadyExist = errors.New("element already exist")

// TasksStorage implements domain.TaskBD
type TasksStorage struct {
	mu     sync.Mutex
	bucket map[domain.TaskID]*domain.Task
}

func NewTasksStorage() domain.TaskDB {
	return &TasksStorage{
		bucket: make(map[domain.TaskID]*domain.Task),
		mu:     sync.Mutex{},
	}
}

func (db *TasksStorage) Create(t *domain.Task) error {
	if _, ok := db.bucket[t.ID]; ok {
		return ErrTaskAlreadyExist
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.bucket[t.ID] = t
	return nil
}

func (db *TasksStorage) Get(id domain.TaskID) (*domain.Task, error) {
	if c, ok := db.bucket[id]; ok {
		return c, nil
	}

	return nil, ErrTaskNotFound
}

func (db *TasksStorage) Update(t *domain.Task) error {
	_, err := db.Get(t.ID)
	if err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.bucket[t.ID] = t
	return nil
}

func (db *TasksStorage) Delete(id domain.TaskID) error {
	_, err := db.Get(id)
	if err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.bucket, id)
	return nil
}

func (db *TasksStorage) List() []*domain.Task {
	x := []*domain.Task{}

	for _, v := range db.bucket {
		x = append(x, v)
	}
	return x
}
