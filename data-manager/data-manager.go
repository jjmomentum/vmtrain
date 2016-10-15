// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
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
	"github.com/vmtrain/data-manager/app"
	"github.com/vmtrain/data-manager/template"
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
	app.Cntxt.Stats.AddHit(r.RequestURI)
	t := template.New(app.Cntxt.ContentRoot, app.Cntxt.APIHost+":"+strconv.Itoa(app.Cntxt.ListenPort))
	t.IndexHandler(w, r)
}

func statsHitsHandler(w http.ResponseWriter, r *http.Request) {
	app.Cntxt.Stats.AddHit(r.RequestURI)
	t := template.New(app.Cntxt.ContentRoot, app.Cntxt.APIHost+":"+strconv.Itoa(app.Cntxt.ListenPort))
	t.StatsHitsHandler(w, r)
}

// Called by main, which is just a wrapper for this function. The reason
// is main can't directly pass back a return code to the OS.
func realMain() int {
	// setup JSON request handlers
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		// stats
		rest.Get("/stats/hits", app.Cntxt.Stats.Get),
		// lab pool data
		rest.Post("/api/reservations", app.CreateReservation),
		rest.Get("/api/reservations", app.ShowReservationList),
		rest.Post("/api/servers", app.CreateServer),
		rest.Get("/api/servers", app.ShowServerList),
		rest.Post("/api/users", app.CreateUser),
		rest.Get("/api/users", app.ShowUserList),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// setup the html page request handlers and mux it all
	mux := http.NewServeMux()
	mux.Handle("/api/", api.MakeHandler())
	mux.Handle("/stats/", api.MakeHandler())
	mux.Handle("/html/skeleton/", http.FileServer(http.Dir(app.Cntxt.ContentRoot)))
	mux.Handle("/html/tmpl/index", http.HandlerFunc(templateHomeHandler))
	mux.Handle("/html/tmpl/hits", http.HandlerFunc(statsHitsHandler))

	// this runs a server that can handle os signals for clean shutdown.
	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(app.Cntxt.ListenPort),
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
