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
	"log"
	"net/http"

	"github.com/vmtrain/approval-monitor/models"
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
func (b Backend) GetServers() (*models.ServerList, int, error) {
	blob, err := b.datastore.Read(blobId)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	serverList := []models.Server{}
	for _, server := range blob.Content.Servers {
		serverList = append(serverList, server)
	}
	var servers models.ServerList = serverList

	return &servers, http.StatusOK, nil
}

// SaveServer is a function to stored the data about a server.
func (b Backend) SaveServer(server models.Server) (*models.Server, int, error) {
	// Read data from the blob service
	blob, err := b.datastore.Read(blobId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Printf("The blob found was %+v", blob)

	blob.Content.Lock()
	blob.Content.Servers[server.UUID] = server
	blob.Content.Unlock()

	err = b.datastore.Write(blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &server, http.StatusOK, nil
}

// GetReservations is a function to look up data about multiple servers.
func (b Backend) GetReservations() (*models.ReservationList, int, error) {
	blob, err := b.datastore.Read(blobId)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	reservationList := []models.Reservation{}
	for _, reservation := range blob.Content.Reservations {
		reservationList = append(reservationList, reservation)
	}
	var reservations models.ReservationList = reservationList

	return &reservations, http.StatusOK, nil
}

// SaveReservation is a function to stored the data about a server.
func (b Backend) SaveReservation(reservation models.Reservation) (*models.Reservation, int, error) {
	// Read data from the blob service
	blob, err := b.datastore.Read(blobId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Printf("The blob found was %+v", blob)

	blob.Content.Lock()
	blob.Content.Reservations[reservation.UUID] = reservation
	blob.Content.Unlock()

	err = b.datastore.Write(blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &reservation, http.StatusOK, nil
}

// GetUsers is a function to look up data about multiple users.
func (b Backend) GetUsers() (*models.UserList, int, error) {
	blob, err := b.datastore.Read(blobId)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	userList := []models.User{}
	for _, user := range blob.Content.Users {
		userList = append(userList, user)
	}
	var users models.UserList = userList

	return &users, http.StatusOK, nil
}

// SaveUser is a function to stored the data about a user.
func (b Backend) SaveUser(user models.User) (*models.User, int, error) {
	blob, err := b.datastore.Read(blobId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Printf("The blob found was %+v", blob)

	blob.Content.Lock()
	blob.Content.Users[user.UUID] = user
	blob.Content.Unlock()

	err = b.datastore.Write(blob)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &user, http.StatusOK, nil
}
