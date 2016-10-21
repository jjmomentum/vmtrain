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

	"github.com/vmtrain/data-manager/models"
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
	content, err := b.datastore.Read()
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	serverList := []models.Server{}
	content.Lock()
	for _, server := range content.Servers {
		serverList = append(serverList, server)
	}
	content.Unlock()
	var servers models.ServerList = serverList

	return &servers, http.StatusOK, nil
}

// SaveServer is a function to stored the data about a server.
func (b Backend) SaveServer(server models.Server) (*models.Server, int, error) {
	// Read data from the blob service
	content, err := b.datastore.Read()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Printf("The content found was %+v", content)

	content.Lock()
	content.Servers[server.UUID] = server
	content.Unlock()

	err = b.datastore.Write(content)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &server, http.StatusOK, nil
}

// DeleteServer is a function to delete the data about a server.
func (b Backend) DeleteServer(id string) (int, error) {
	// Read data from the blob service
	content, err := b.datastore.Read()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	log.Printf("The content found was %+v\n", content)

	content.Lock()
	_, ok := content.Servers[id]
	if !ok {
		content.Unlock()
		return http.StatusNotFound, fmt.Errorf("Server %s was not found", id)
	} else {
		delete(content.Servers, id)
		content.Unlock()
	}

	err = b.datastore.Write(content)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// GetReservations is a function to look up data about multiple servers.
func (b Backend) GetReservations() (*models.ReservationList, int, error) {
	content, err := b.datastore.Read()
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	reservationList := []models.Reservation{}
	content.Lock()
	for _, reservation := range content.Reservations {
		reservationList = append(reservationList, reservation)
	}
	content.Unlock()
	var reservations models.ReservationList = reservationList

	return &reservations, http.StatusOK, nil
}

// GetReservation is a function to look up data about a single reservation.
func (b Backend) GetReservation(id string) (*models.Reservation, int, error) {
	content, err := b.datastore.Read()
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}
	content.Lock()
	res, ok := content.Reservations[id]
	content.Unlock()
	if !ok {
		return nil,
			http.StatusNotFound,
			fmt.Errorf("Reservation with uuid: %s not found", id)
	}

	return &res, http.StatusOK, nil
}

// SaveReservation is a function to stored the data about a server.
func (b Backend) SaveReservation(reservation models.Reservation) (*models.Reservation, int, error) {
	var returnReservation *models.Reservation
	// Read data from the blob service
	content, err := b.datastore.Read()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Printf("The content found was %+v", content)

	// Update or create new reservation in the blob
	content.Lock()
	resFound, ok := content.Reservations[reservation.UUID]
	if ok {
		resFound.Update(reservation)
		content.Reservations[reservation.UUID] = resFound
		returnReservation = &resFound
	} else {
		content.Reservations[reservation.UUID] = reservation
		returnReservation = &reservation
	}
	content.Unlock()

	err = b.datastore.Write(content)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return returnReservation, http.StatusOK, nil
}

// GetUsers is a function to look up data about multiple users.
func (b Backend) GetUsers() (*models.UserList, int, error) {
	content, err := b.datastore.Read()
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Failed to read blob data from service. Caused by: %v", err)
	}

	userList := []models.User{}
	content.Lock()
	for _, user := range content.Users {
		userList = append(userList, user)
	}
	content.Unlock()
	var users models.UserList = userList

	return &users, http.StatusOK, nil
}

// SaveUser is a function to stored the data about a user.
func (b Backend) SaveUser(user models.User) (*models.User, int, error) {
	content, err := b.datastore.Read()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Printf("The content found was %+v", content)

	content.Lock()
	content.Users[user.UUID] = user
	content.Unlock()

	err = b.datastore.Write(content)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &user, http.StatusOK, nil
}

// DeleteUser is a function to delete the data about a user.
func (b Backend) DeleteUser(id string) (int, error) {
	content, err := b.datastore.Read()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	log.Printf("The content found was %+v", content)

	content.Lock()
	_, ok := content.Users[id]
	if !ok {
		content.Unlock()
		return http.StatusNotFound, fmt.Errorf("User %s was not found", id)
	} else {
		delete(content.Users, id)
		content.Unlock()
	}

	err = b.datastore.Write(content)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
