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

import "github.com/maantos/todoApplication/pkg/data"

//
// NOTE: Types that are defined here are used purely for documentation purpose,.
// None of the handlers is using them

// A list of TODO tasks
//
//swagger:response tasksResponse
type tasksResponseWrapper struct {
	// All current TODO tasks
	// in: body
	Body []data.Task
}
