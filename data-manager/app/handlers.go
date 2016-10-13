// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import "net/http"

// CreateServer creates a job and stores it in etcd
func CreateServer(w http.ResponseWriter, r *http.Request) {
	//TODO add logic to create a server and store it in the blob service
}

// ShowServerList displays a list of jobs stored in etcd
func ShowServerList(w http.ResponseWriter, r *http.Request) {
	// TODO add logic to show a list of reservations
}

// CreateReservation creates a job and stores it in etcd
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	// TODO add logic to create a reservation and add it to the blob service
}

// ShowReservationList displays a list of jobs stored in etcd
func ShowReservationList(w http.ResponseWriter, r *http.Request) {
	// TODO add logic to show a list of reservations
}
