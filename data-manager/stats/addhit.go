// Package stats collects statistics about the application.
//
// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package stats

import (
	"log"
)

// AddHit adds a hit
func (s *Stats) AddHit(request string) {
	s.lock.Lock()
	count := s.Hits[request]
	s.Hits[request] = count + 1
	log.Printf("Counting hit: %s -- up to %d\n", request, count)
	s.lock.Unlock()
}
