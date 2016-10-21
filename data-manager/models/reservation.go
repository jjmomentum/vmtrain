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
	"log"
)

// Reservation is the struct to hold reservation data about the lab environment
type Reservation struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	ServerName string `json:"server_name"`

	Approved bool `json:"approved"`
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

// Update changes the fields in current model to the ones of the model passed in.
func (r *Reservation) Update(res Reservation) {
	log.Printf("Updating reservation %+v with %+v", r, res)
	if res.EndDate != "" {
		r.EndDate = res.EndDate
	}

	if res.Name != "" {
		r.Name = res.Name
	}

	if res.ServerName != "" {
		r.ServerName = res.ServerName
	}

	if res.StartDate != "" {
		r.StartDate = res.StartDate
	}

	r.Approved = res.Approved
}

// ReservationList is a list of Reservation structs
type ReservationList []Reservation
