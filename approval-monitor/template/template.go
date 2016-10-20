// Package template holds templates for all the packages.
//
// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package template

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vmtrain/approval-monitor/stats"
)

// Template is a client to get topics
type Template struct {
	ContentRoot string
	APIHost     string
}

// New returns a new Template object initialized -- convenience function.
func New(contentRoot string, apiHost string) Template {
	return Template{
		ContentRoot: contentRoot,
		APIHost:     apiHost,
	}
}

func init() {
	log.Println("Initialized Template.")
}

func (t *Template) generateAPIUrl(path string) string {
	return "http://" + t.APIHost + path
}

// Retrieve stats info via REST call.
func (t *Template) getStatsHits() map[string]int {
	url := t.generateAPIUrl("/stats/hits")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := stats.HitsFromJSON(body)
	perror(err)

	return data
}
