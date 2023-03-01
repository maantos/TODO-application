package db

import (
	"reflect"
	"testing"

	"github.com/maantos/todoApplication/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateListTasks(t *testing.T) {
	ts := NewTasksStorage()

	var task = &domain.Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}

	err := ts.Create(task)
	assert.NoError(t, err, "=> Create failed")

	want := []*domain.Task{task}

	if got := ts.List(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

func TestCreateTasksThatExist(t *testing.T) {
	ts := NewTasksStorage()
	var task = &domain.Task{ID: "1", Done: true, Title: "test Title", Description: "Todo task 1."}

	err := ts.Create(task)
	assert.NoError(t, err, "=> Create failed")

	err = ts.Create(task)
	assert.Error(t, err)
	assert.Equal(t, err, ErrTaskAlreadyExist)
}

func TestCreateUpdateTasks(t *testing.T) {
	ts := NewTasksStorage()

	var task = &domain.Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}

	err := ts.Create(task)
	assert.NoError(t, err, "=> Create failed")

	t.Run("Update task that doesnt exist", func(t *testing.T) {
		var updated_task = &domain.Task{ID: "12", Done: true, Title: "New task", Description: "Todo task 12."}

		err := ts.Update(updated_task)
		assert.Error(t, err)
		assert.Equal(t, err, ErrTaskNotFound)
	})

	t.Run("Update task", func(t *testing.T) {
		var updated_task = &domain.Task{ID: "1", Done: true, Title: "Task after update", Description: "Todo task 1."}

		err := ts.Update(updated_task)
		assert.NoError(t, err, "=> update failed")

		get, err := ts.Get(updated_task.ID)
		assert.Nil(t, err)

		if !reflect.DeepEqual(get, updated_task) {
			t.Errorf("=> failed to update task, got %v wanted %v", get, updated_task)
		}
	})
}

func TestAddReadTODOTasks(t *testing.T) {
	t.Log("Delete Todo tasks")

	ts := NewTasksStorage()

	t.Run("Get element from empty store", func(t *testing.T) {
		_, err := ts.Get("1")

		assert.Error(t, err)
		assert.Equal(t, err, ErrTaskNotFound)
	})

	var task = &domain.Task{ID: "1", Done: false, Title: "test Title", Description: "Todo task 1."}
	t.Run("Get single element from the store", func(t *testing.T) {
		err := ts.Create(task)
		assert.NoError(t, err, "=> Create failed")

		var got *domain.Task
		got, err = ts.Get(task.ID)
		assert.NoError(t, err, "=> failed to read task")

		if !reflect.DeepEqual(got, task) {
			t.Errorf("=> got %v wanted %v", got, task)
		}
	})

}

func TestDeleteTODOTasks(t *testing.T) {
	t.Log("Delete Todo tasks")

	ts := NewTasksStorage()

	t.Run("Delete element that doesnt exist", func(t *testing.T) {
		err := ts.Delete("1")

		assert.Error(t, err)
		assert.Equal(t, err, ErrTaskNotFound)

	})

	var task = &domain.Task{ID: "13", Done: false, Title: "To be Deleted", Description: "Task to be removed."}
	t.Run("Delete element from the store", func(t *testing.T) {
		err := ts.Create(task)
		assert.NoError(t, err, "=> Create failed")

		err = ts.Delete("13")
		assert.NoError(t, err, "=> deletion failed")
	})
}
