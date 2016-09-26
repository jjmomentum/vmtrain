// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	html_template "html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tdhite/q3-training-journal/journal"
)

type Topics struct {
	Topic    string
	Messages []journal.Message
}

// Generate topic Message list.
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
