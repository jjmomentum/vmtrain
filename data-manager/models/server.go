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

// Server is the struct to hold server data about the lab environment
type Server struct {
	Name int `json:"name"`
}

// FromJSON populates Server from JSON data.
func (s *Server) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, s)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the Server.
func (s *Server) ToJSON() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ServerList is a list of Server models
type ServerList struct {
	Servers []Server `json:"servers"`
}

// FromJSON populates ServerList from JSON data.
func (sl *ServerList) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, sl)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the ServerList.
func (sl *ServerList) ToJSON() (string, error) {
	b, err := json.Marshal(sl)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
