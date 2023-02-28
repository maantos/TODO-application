package app

import (
	"github.com/maantos/todoApplication/pkg/domain"
)

// taskService implements domain.TaskService
type taskService struct {
	db domain.TaskDB
}

func NewTaskService(db domain.TaskDB) domain.TaskService {
	return &taskService{
		db: db,
	}
}

func (ts *taskService) Get(id domain.TaskID) (*domain.Task, error) {
	return ts.db.Get(id)

}

func (ts *taskService) List() []*domain.Task {
	return ts.db.List()
}

func (ts *taskService) Create(task *domain.Task) error {
	return ts.db.Create(task)
}

func (ts *taskService) Update(t *domain.Task) error {
	return ts.db.Update(t)
}

func (ts *taskService) Delete(id domain.TaskID) error {
	return ts.db.Delete(id)
}
