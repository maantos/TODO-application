package domain

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

type TaskDB interface {
	Create(*Task) error
	Get(TaskID) (*Task, error)
	Update(*Task) error
	Delete(TaskID) error
	List() []*Task
}
