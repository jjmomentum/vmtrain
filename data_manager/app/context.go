// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT 
//
package app

import "github.com/tdhite/q3-training-journal/stats"

// Global application context variables.
type AppContext struct {
	ListenPort  int
	ContentRoot string
	APIHost     string
	Stats       stats.Stats
}

func New() *AppContext {
	ctx := &AppContext{
		ListenPort:  80,
		ContentRoot: ".",
		APIHost:     "localhost",
		Stats:       stats.New(),
	}
	return ctx
}
