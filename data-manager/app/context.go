// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"github.com/vmtrain/data-manager/models"
	"github.com/vmtrain/data-manager/stats"
)

// Context is a struct to hold global application context variables.
type Context struct {
	ListenPort  int
	ContentRoot string
	APIHost     string
	Stats       stats.Stats
	Backend     Backend
}

// New generates an AppContext struct
func New() *Context {
	content := models.Content{
		Servers:      map[string]models.Server{},
		Reservations: map[string]models.Reservation{},
		Users:        map[string]models.User{},
	}
	contentJSON, err := content.ToJSON()
	if err != nil {
		panic("Failed to marshal the content of the blob into JSON")
	}
	// TODO Remove once we start integration with the blob service
	bcknd := NewBackend(
		NewMockDatastore(
			map[int]*models.Blob{
				blobId: &models.Blob{
					ID:      blobId,
					Name:    "Team 4 blob",
					Content: contentJSON,
				},
			},
		),
	)
	ctx := &Context{
		ListenPort:  80,
		ContentRoot: ".",
		APIHost:     "localhost",
		Stats:       stats.New(),
		Backend:     bcknd,
	}
	return ctx
}
