package data

import (
	"reflect"
	"testing"
)

func TestCreateListTasks(t *testing.T) {
	t.Log("Create and List todo tasks")

	ts := NewTasksDB()
	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}

	err := ts.Create(task)

	if err != nil {
		t.Errorf("=> Create failed with %v", err.Error())
	}

	want := []*Task{task}

	if got := ts.List(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

func TestCreateTasksThatExist(t *testing.T) {
	t.Log("Create Todo tasks that already exists")

	ts := NewTasksDB()
	var task = &Task{ID: "1", Done: true, Title: "test Title", Description: "Todo task 1."}

	err := ts.Create(task)

	if err != nil {
		t.Errorf("=> Create failed with %v", err.Error())
	}

	err = ts.Create(task)

	if err != ErrTaskAlreadyExist {
		t.Error("=> Creating the same task should return an error!")
	}
}

func TestCreateUpdateTasks(t *testing.T) {
	t.Log("Create and Update task")

	ts := NewTasksDB()

	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	err := ts.Create(task)

	if err != nil {
		t.Errorf("=> failed to create task: %v", err.Error())
	}

	t.Run("Update task that doenst exist", func(t *testing.T) {
		var updated_task = &Task{ID: "12", Done: true, Title: "New task", Description: "Todo task 12."}
		err := ts.Update(updated_task)

		if err != ErrTaskNotFound {
			t.Errorf("=> updating task that doesnt exist should return error")
		}
	})

	t.Run("Update task", func(t *testing.T) {
		var updated_task = &Task{ID: "1", Done: true, Title: "Task after update", Description: "Todo task 1."}

		err := ts.Update(updated_task)

		if err != nil {
			t.Errorf("=> failed to update task: %v", err.Error())
		}

		if get, err := ts.Read(updated_task.ID); !reflect.DeepEqual(get, updated_task) || err != nil {
			t.Errorf("=> failed to update task, got %v wanted %v", get, updated_task)
		}
	})
}

func TestAddReadTODOTasks(t *testing.T) {
	t.Log("Delete Todo tasks")

	ts := NewTasksDB()

	t.Run("Get element from empty store", func(t *testing.T) {
		elem, err := ts.Read("1")
		if elem != nil || err != ErrTaskNotFound {
			t.Errorf("=> Reading task that doesn't exist should failed.")
		}
	})

	var task = &Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	t.Run("Get single element from the store", func(t *testing.T) {
		err := ts.Create(task)
		if err != nil {
			t.Errorf("=> failed to create task: %v", err.Error())
		}

		var got *Task
		got, err = ts.Read(task.ID)
		if err != nil {
			t.Errorf("=> failed to read task: %v", err.Error())
		}

		if !reflect.DeepEqual(got, task) {
			t.Errorf("=> got %v wanted %v", got, task)
		}
	})

}

func TestDeleteTODOTasks(t *testing.T) {
	t.Log("Delete Todo tasks")

	ts := NewTasksDB()

	t.Run("Delete element that doesnt exist", func(t *testing.T) {
		err := ts.Delete("1")

		if err != ErrTaskNotFound {
			t.Error("=> Missing element removal should throw an error")
		}

	})

	var task = &Task{ID: "13", Done: false, Title: "To be Deleted", Description: "Task to be removed."}
	t.Run("Delete element from the store", func(t *testing.T) {
		err := ts.Create(task)

		if err != nil {
			t.Errorf("=> failed to create task: %v", err.Error())
		}

		err = ts.Delete("13")

		if err != nil {
			t.Errorf("=> failed to delete element: %v", err.Error())
		}

	})
}
