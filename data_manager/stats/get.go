// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT 
//
package stats

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Stats) Get(w rest.ResponseWriter, r *rest.Request) {
	s.AddHit(r.RequestURI)
	s.lock.Lock()
	err := w.WriteJson(s.hits)
	s.lock.Unlock()
	if err != nil {
		rest.Error(w, err.Error(), 503)
	}
}
