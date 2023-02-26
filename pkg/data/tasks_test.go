package data

import (
	"reflect"
	"testing"
)

type StubTaskStore struct {
	data map[TaskID]*Task
}

func (st *StubTaskStore) Create(t *Task) error {
	if _, ok := st.data[t.ID]; ok {
		return ErrTaskAlreadyExist
	}
	st.data[t.ID] = t
	return nil
}
func (s *StubTaskStore) Read(id TaskID) (*Task, error) {
	if c, ok := s.data[id]; ok {
		return c, nil
	}
	return nil, ErrTaskNotFound
}
func (s *StubTaskStore) Update(t *Task) error {
	_, err := s.Read(t.ID)
	if err != nil {
		return err
	}
	s.data[t.ID] = t
	return nil

}
func (s *StubTaskStore) Delete(id TaskID) error {
	_, err := s.Read(id)
	if err != nil {
		return err
	}
	delete(s.data, id)
	return nil

}
func (s *StubTaskStore) List() []*Task {
	x := []*Task{}

	for _, v := range s.data {
		x = append(x, v)
	}
	return x

}

func TestListTODOTasks(t *testing.T) {
	t.Log("List Todo tasks")

	ts := NewTasksDB()

	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	ts.Create(task)

	want := []*Task{task}

	if got := ts.List(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

func TestCreateTODOTasks(t *testing.T) {
	t.Log("Create Todo tasks")

	var stubTask = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	store := StubTaskStore{
		map[TaskID]*Task{
			"1": stubTask,
		},
	}
	t.Run("Create new task", func(t *testing.T) {
		var newTask = &Task{ID: "2", Done: false, Title: "New task", Description: "Todo task 2."}
		err := store.Create(newTask)
		if err != nil {
			t.Errorf("=> Create failed with %v", err.Error())
		}
	})

	t.Run("Create task, that already exist", func(t *testing.T) {
		err := store.Create(stubTask)
		if err != ErrTaskAlreadyExist {
			t.Errorf("=> Create failed with %v", err.Error())
		}
	})
}

func TestUpdateTODOTasks(t *testing.T) {
	t.Log("Update Todo tasks")

	ts := NewTasksDB()
	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	err := ts.Create(task)

	if err != nil {
		t.Errorf("=> failed to create task: %v", err.Error())
	}
	var updated_task = &Task{ID: "1", Done: true, Title: "Updated test Title", Description: "Todo task 1."}

	err = ts.Update(updated_task)

	if err != nil {
		t.Errorf("=> failed to update task: %v", err.Error())
	}

}

func TestReadTODOTasks(t *testing.T) {
	t.Log("Delete Todo tasks")

	ts := NewTasksDB()

	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	err := ts.Create(task)
	if err != nil {
		t.Errorf("=> failed to create task: %v", err.Error())
	}
	var got *Task
	got, err = ts.Read(task.ID)
	if err != nil {
		t.Errorf("=> failed to read task: %v", err.Error())
	}

	if got != task {
		t.Errorf("Got %v wanted %v", got, task)
	}

}

func TestDeleteTODOTasks(t *testing.T) {
	t.Log("Delete Todo tasks")

	ts := NewTasksDB()

	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	err := ts.Create(task)

	if err != nil {
		t.Errorf("=> Create failed with %v", err.Error())
	}

	err = ts.Delete(task.ID)
	if err != nil {
		t.Errorf("=> failed to delete task %v", err.Error())
	}
}
