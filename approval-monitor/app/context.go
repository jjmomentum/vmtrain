// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"github.com/vmtrain/approval-monitor/models"
	"github.com/vmtrain/approval-monitor/stats"
)

// Context is a struct to hold global application context variables.
type Context struct {
	ListenPort  int
	ContentRoot string
	APIHost     string
	Stats       stats.Stats
	backend     Backend
}

// New generates an AppContext struct
func New() *Context {
	// TODO Remove once we start integration with the blob service
	bcknd := NewBackend(
		NewMockDatastore(
			map[int]*models.Blob{
				blobId: &models.Blob{
					ID:   blobId,
					Name: "Team 4 blob",
					Content: models.Content{
						Servers:      map[string]models.Server{},
						Reservations: map[string]models.Reservation{},
						Users:        map[string]models.User{},
					},
				},
			},
		),
	)
	ctx := &Context{
		ListenPort:  80,
		ContentRoot: ".",
		APIHost:     "localhost",
		Stats:       stats.New(),
		backend:     bcknd,
	}
	return ctx
}
