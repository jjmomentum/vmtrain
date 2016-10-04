// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/stretchr/graceful"
	"github.com/tdhite/q3-training-journal/app"
	"github.com/tdhite/q3-training-journal/journal"
	"github.com/tdhite/q3-training-journal/template"
)

// Http handler functions for dealing with various site requests for
// home page, editing, deleting and saving objects.
//
// These are not all that necessary as they are just a trick to use the
// http.ServeMux to create a poor man's URL router. The json stuff uses
// the venerable go-json-router, but the site pages are so simple it's not
// worth writing up a whole router model just for that when we can just 'mux'
// things via separate handlers for each html (site) request.
func templateHomeHandler(w http.ResponseWriter, r *http.Request) {
	app.Context.Stats.AddHit(r.RequestURI)
	t := template.New(app.Context.ContentRoot, app.Context.APIHost+":"+strconv.Itoa(app.Context.ListenPort))
	t.IndexHandler(w, r)
}

func templateTopicHandler(w http.ResponseWriter, r *http.Request) {
	app.Context.Stats.AddHit(r.RequestURI)
	t := template.New(app.Context.ContentRoot, app.Context.APIHost+":"+strconv.Itoa(app.Context.ListenPort))
	t.TopicHandler(w, r)
}

func statsHitsHandler(w http.ResponseWriter, r *http.Request) {
	app.Context.Stats.AddHit(r.RequestURI)
	t := template.New(app.Context.ContentRoot, app.Context.APIHost+":"+strconv.Itoa(app.Context.ListenPort))
	t.StatsHitsHandler(w, r)
}

// Called by main, which is just a wrapper for this function. The reason
// is main can't directly pass back a return code to the OS.
func realMain() int {
	// setup JSON request handlers
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	jrnl := journal.NewMock()

	router, err := rest.MakeRouter(
		// stats
		rest.Get("/stats/hits", app.Context.Stats.Get),
		// topics
		rest.Get("/api/topics", jrnl.RestGetTopics),
		rest.Post("/api/topic/:topic", jrnl.RestPostTopic),
		rest.Get("/api/topic/:topic", jrnl.RestGetTopic),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// setup the html page request handlers and mux it all
	mux := http.NewServeMux()
	mux.Handle("/api/", api.MakeHandler())
	mux.Handle("/stats/", api.MakeHandler())
	mux.Handle("/html/skeleton/", http.FileServer(http.Dir(app.Context.ContentRoot)))
	mux.Handle("/html/tmpl/index", http.HandlerFunc(templateHomeHandler))
	mux.Handle("/html/tmpl/topic", http.HandlerFunc(templateTopicHandler))
	mux.Handle("/html/tmpl/hits", http.HandlerFunc(statsHitsHandler))

	// this runs a server that can handle os signals for clean shutdown.
	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(app.Context.ListenPort),
			Handler: mux,
		},
		ListenLimit: 1024,
	}

	exitcode := 0
	err = server.ListenAndServe()
	if err != nil {
		log.Println("Shutdown caused by:" + err.Error())
		exitcode = 1
	}

	return exitcode
}

func main() {
	// Delegate to realMain so defered operations can happen (os.Exit exits
	// the program without servicing defer statements)
	os.Exit(realMain())
}
