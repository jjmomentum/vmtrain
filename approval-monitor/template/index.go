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
)

// IndexHandler generates the main (home) page of the site.
func (t *Template) IndexHandler(w http.ResponseWriter, r *http.Request) {

	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		if err := tmpl.ExecuteTemplate(w, page, map[string]string{"hello": "world"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
