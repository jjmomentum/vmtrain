// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pborman/uuid"
	"github.com/vmtrain/data-manager/models"
)

// CreateServer creates a server
func CreateServer(w rest.ResponseWriter, r *rest.Request) {
	var server models.Server
	// Read and validate the request. The read on the request body is limited
	// to prevent malicious attacks on the server.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		defer r.Body.Close()
		log.Printf("The body received is %s", string(body))

		// Unmarshal and validate JSON
		err = json.Unmarshal(body, server)
		if err != nil {
			rest.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			if server.Name == "" {
				rest.Error(w, "Missing 'name' in the request payload", http.StatusBadRequest)
			} else {
				// Store in the blob service
				savedServer, status, err := Cntxt.backend.SaveServer(server)
				if err != nil {
					rest.Error(w, err.Error(), status)
				} else {
					w.WriteJson(savedServer)
				}
			}
		}
	}
}

// ShowServerList displays a list of servers
func ShowServerList(w rest.ResponseWriter, r *rest.Request) {
	serverList, status, err := Cntxt.backend.GetServers()
	if err != nil {
		rest.Error(w, err.Error(), status)
	} else {
		w.WriteJson(serverList)
	}
}

// CreateReservation creates a reservation
func CreateReservation(w rest.ResponseWriter, r *rest.Request) {
	var reservation models.Reservation
	// Read and validate the request. The read on the request body is limited
	// to prevent malicious attacks on the server.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		defer r.Body.Close()
		log.Printf("The body received is %s", string(body))

		// Unmarshal and validate JSON
		err = json.Unmarshal(body, reservation)
		if err != nil {
			rest.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			if reservation.UUID == "" {
				reservation.UUID = uuid.New()
			}

			// Store in the blob service
			savedReservation, status, err := Cntxt.backend.SaveReservation(reservation)
			if err != nil {
				rest.Error(w, err.Error(), status)
			} else {
				w.WriteJson(savedReservation)
			}

		}
	}
}

// ShowReservationList displays a list of reservations
func ShowReservationList(w rest.ResponseWriter, r *rest.Request) {
	reservationList, status, err := Cntxt.backend.GetReservations()
	if err != nil {
		rest.Error(w, err.Error(), status)
	} else {
		w.WriteJson(reservationList)
	}
}

// CreateUser creates a user
func CreateUser(w rest.ResponseWriter, r *rest.Request) {
	var user models.User
	// Read and validate the request. The read on the request body is limited
	// to prevent malicious attacks on the server.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		defer r.Body.Close()
		log.Printf("The body received is %s", string(body))

		// Unmarshal and validate JSON
		err = json.Unmarshal(body, user)
		if err != nil {
			rest.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			if user.UUID == "" {
				user.UUID = uuid.New()
			}

			// Store in the blob service
			savedUser, status, err := Cntxt.backend.SaveUser(user)
			if err != nil {
				rest.Error(w, err.Error(), status)
			} else {
				w.WriteJson(savedUser)
			}

		}
	}
}

// ShowUserList displays a list of users
func ShowUserList(w rest.ResponseWriter, r *rest.Request) {
	userList, status, err := Cntxt.backend.GetUsers()
	if err != nil {
		rest.Error(w, err.Error(), status)
	} else {
		w.WriteJson(userList)
	}
}
