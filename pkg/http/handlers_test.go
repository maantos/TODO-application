package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maantos/todoApplication/pkg/data"
)

type StubTaskStore struct {
	TaskFunc func(t *data.Task) error
	ReadFunc func(id data.TaskID) (*data.Task, error)
}

func (s *StubTaskStore) Create(t *data.Task) error {
	if s.TaskFunc != nil {
		return s.TaskFunc(t)
	}
	return nil
}

func (s *StubTaskStore) Read(id data.TaskID) (*data.Task, error) {
	if s.ReadFunc != nil {
		s.ReadFunc(id)
	}

	if c, ok := db.bucket[id]; ok {
		return c, nil
	}

	return nil, data.ErrTaskNotFound
}

func (db *StubTaskStore) Update(t *data.Task) error {
	_, err := db.Read(t.ID)
	if err != nil {
		return err
	}

	db.bucket[t.ID] = t
	return nil
}

func (db *StubTaskStore) Delete(id data.TaskID) error {
	_, err := db.Read(id)
	if err != nil {
		return err
	}

	delete(db.bucket, id)
	return nil
}

func (db *StubTaskStore) List() []*data.Task {
	x := []*data.Task{}

	for _, v := range db.bucket {
		x = append(x, v)
	}
	return x
}

func TestGetTasks(t *testing.T) {
	t.Run("Return TODO tasks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		response := httptest.NewRecorder()

		handler := &Tasks{
			db: &StubTaskStore{},
		}
		handler.ListAll(response, request)
		got := response.Body.String()
		want := "[]data.Task{}"

		if got != want {
			t.Errorf("=>got %v wanted %v", got, want)
		}

	})
}
