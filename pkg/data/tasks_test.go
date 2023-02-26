package data

import (
	"reflect"
	"testing"
)

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

	ts := NewTasksDB()

	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	err := ts.Create(task)

	if err != nil {
		t.Errorf("=> Create failed with %v", err.Error())
	}
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
