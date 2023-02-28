// Package classification TODO Tasks API.
//
// The purpose of this application is to provide an application
// that is using plain go code to define an API
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http
//	Host: localhost
//	BasePath: /v1
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/maantos/todoApplication/pkg/domain"

//
// NOTE: Types that are defined here are used purely for documentation purpose,.
// None of the handlers is using them

// A list of todo-tasks
//
//swagger:response tasksResponse
type tasksResponseWrapper struct {
	// All current TODO tasks
	// in: body
	Body []*domain.Task
}

// Single TODO task
//
//swagger:response taskResponse
type taskResponseWrapper struct {
	// Single todo-task
	// in: body
	Body domain.Task
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// Error message descriptions
	ErrorMessage string
}

// swagger:response createdTask
type taskIDResponsewrapper struct {
	// The id of the task for which the operation relates
	ID string `json:"id"`
}

// swagger:parameters deleteTask
type taskIDParamsWrapper struct {
	// The id of the task for which the operation relates
	ID string `json:"id"`
}

// swagger:parameters createTask
type taskParamsWrapper struct {
	// Product data structure to Update or Create.
	// in: body
	// required: true
	Body domain.Task
}
