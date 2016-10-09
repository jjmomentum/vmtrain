// Package stats collects statistics about the application.
//
// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package stats

import (
	"encoding/json"
	"log"
	"sync"
)

// Stats store hits per URL
type Stats struct {
	Hits map[string]int `json:"hits"`
	lock *sync.RWMutex
}

func init() {
	log.Println("Initialized stats package.")
}

// New returns a new Stats struct
func New() Stats {
	return Stats{
		Hits: make(map[string]int),
		lock: &sync.RWMutex{},
	}
}

// HitsFromJSON converts a JSON string to Go struct and return.
func HitsFromJSON(jsonData []byte) (map[string]int, error) {
	var hits map[string]int
	err := json.Unmarshal([]byte(jsonData), &hits)
	if err != nil {
		log.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			log.Println(string(jsonData[v.Offset-40 : v.Offset]))
		}
	}

	return hits, err
}
