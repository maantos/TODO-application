package handlers

import (
	"log"

	"github.com/maantos/todoApplication/pkg/data"
)

// struct used as key in context
type TaskKey struct{}

type KeyProduct struct{}

type Tasks struct {
	l  *log.Logger
	db *data.TasksDB
}

func NewTasksHandler(logger *log.Logger, db *data.TasksDB) *Tasks {
	return &Tasks{
		l:  logger,
		db: db,
	}
}