// Package template holds templates for all the packages.
//
// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package template

import (
	html_template "html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tdhite/q3-training-journal/journal"
)

// Topics holds topics information
type Topics struct {
	Topic    string
	Messages []journal.Message
}

// TopicHandler generates topic Message list.
func (t *Template) TopicHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	topic := r.URL.Query().Get("topic")
	msgs := t.getTopic(topic)
	data := Topics{
		Topic:    topic,
		Messages: msgs,
	}

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		if err := tmpl.ExecuteTemplate(w, page, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
