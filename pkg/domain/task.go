package domain

type TaskID string

//pacakge domain contains all the informations about
// all the objcect and describe interactions between them.
// This package should have no dependencies to any other project
//hexagonal architecture
// Ports -> interfaces
// Adaptors components implementing the contract
//All the subpackages only comunicate by the interfaces defined here.

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

type TaskService interface {
	Create(*Task) error
	Get(TaskID) (*Task, error)
	Update(*Task) error
	Delete(TaskID) error
	List() []*Task
}

type TaskDB interface {
	Create(*Task) error
	Get(TaskID) (*Task, error)
	Update(*Task) error
	Delete(TaskID) error
	List() []*Task
}
