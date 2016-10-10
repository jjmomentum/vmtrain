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
	ID int `json:"id"`

	User string `json:"user"`

	Date string `json:"date"`

	Server string `json:"server"`

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

// ReservationList is a list of Reservation
type ReservationList struct {
	Reservations []Reservation `json:"reservations"`
}

// FromJSON populates ReservationList from JSON data.
func (r *ReservationList) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, r)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the ReservationList.
func (r *ReservationList) ToJSON() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
