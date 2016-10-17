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

// User is the struct to hold user data about the lab environment
type User struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// FromJSON populates User from JSON data.
func (u *User) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, u)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the User.
func (u *User) ToJSON() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// UserList is a list of User structs
type UserList []User
