// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

// Context is a struct to hold global application context variables.
type Context struct {
	DataManagerURL          string
	ApprovalPollIntervalSec int
}

// New generates an AppContext struct
func New() *Context {
	return &Context{}
}
