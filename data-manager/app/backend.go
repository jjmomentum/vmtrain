// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"fmt"
	"net/http"

	"github.com/vmtrain/data-manager/models"
)

// Constants used to map the queries to tables and columns found in the
// database schema
const (
	blobId = 04
)

// Backend is a struct used for storage by the API server. It has a single
// field which is an interface for the storage medium to use via db or in memory
// map as the one used for testing.
type Backend struct {
	datastore Datastore
}

// NewBackend is a constructor that returns a new instance of the Backend
// struct.
func NewBackend(ds Datastore) Backend {
	return Backend{
		datastore: ds,
	}
}

// GetServers is a function to look up data about multiple servers.
func (b Backend) GetServers() (map[string]models.Server, int, error) {
	var (
		err error
	)

	blob := models.Blob{
		ID: blobId,
	}

	servers, err := b.getServersFromBlob(&blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return servers, http.StatusOK, nil
}

// SaveServer is a function to stored the data about a server.
func (b Backend) SaveServer(server models.Server) (*models.Server, int, error) {
	// Read data from the blob service
	blob := models.Blob{
		ID: blobId,
	}
	blob.Content.Servers[server.Name] = server

	err := b.datastore.Write(&blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &server, http.StatusOK, nil
}

// GetReservations is a function to look up data about multiple servers.
func (b Backend) GetReservations() (map[string]models.Reservation, int, error) {
	var (
		err error
	)

	blob := models.Blob{
		ID: blobId,
	}

	reservations, err := b.getReservationsFromBlob(&blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return reservations, http.StatusOK, nil
}

// SaveReservation is a function to stored the data about a server.
func (b Backend) SaveReservation(reservation models.Reservation) (*models.Reservation, int, error) {
	// Read data from the blob service
	blob := models.Blob{
		ID: blobId,
	}

	blob.Content.Reservations[reservation.UUID] = reservation

	err := b.datastore.Write(&blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &reservation, http.StatusOK, nil
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS BELOW
///////////////////////////////////////////////////////////////////////////////

func (b Backend) getServersFromBlob(blob *models.Blob) (map[string]models.Server, error) {
	err := b.datastore.Read(blob)
	if err != nil {
		return nil,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	return blob.Content.Servers, nil
}

func (b Backend) getReservationsFromBlob(blob *models.Blob) (map[string]models.Reservation, error) {
	err := b.datastore.Read(blob)
	if err != nil {
		return nil,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	return blob.Content.Reservations, nil
}
