// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import "github.com/tdhite/q3-training-journal/stats"

// Context is a struct to hold global application context variables.
type Context struct {
	ListenPort  int
	ContentRoot string
	APIHost     string
	Stats       stats.Stats
}

// New generates an AppContext struct
func New() *Context {
	ctx := &Context{
		ListenPort:  80,
		ContentRoot: ".",
		APIHost:     "localhost",
		Stats:       stats.New(),
	}
	return ctx
}
