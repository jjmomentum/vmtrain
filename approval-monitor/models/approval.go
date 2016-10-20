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

// Approval is the struct to hold approval data about the lab environment
type Approval struct {
	ID string `json:"id"`

	TeamID string `json:"teamID"`

	Blob string `json:"blob"`

	Description string `json:"description"`

	Approved bool `json:"approved"`
}

// FromJSON populates Approval from JSON data.
func (a *Approval) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the Approval.
func (a *Approval) ToJSON() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ApprovalList is a list of Approval structs
type ApprovalList []Approval
