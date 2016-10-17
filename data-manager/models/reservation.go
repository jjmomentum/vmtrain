// Package models has the structs that are used by the application.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package models

import (
	"encoding/json"
)

// Reservation is the struct to hold reservation data about the lab environment
type Reservation struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	ServerName string `json:"server_name"`

	Status string `json:"status"`
}

// FromJSON populates Reservation from JSON data.
func (r *Reservation) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, r)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the Reservation.
func (r *Reservation) ToJSON() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReservationList is a list of Reservation structs
type ReservationList struct {
	Reservations []Reservation `json:"reservations"`
}
