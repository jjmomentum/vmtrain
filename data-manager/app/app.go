// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"flag"
	"log"
)

// Cntxt is ues to initialize the state of the service including default ports and host ip
var Cntxt = New()

// Initialize the flags processor with default values and help messages.
func initFlags() {
	const (
		listenPortDefault  = 8080
		listenPortUsage    = "port on which to listen for HTTP requests"
		contentRootDefault = "."
		contentRootUsage   = "path to (static content) templates, skeleton, etc."
		apiHostDefault     = "localhost"
		apiHostUsage       = "host to use for all API calls made internally."
	)

	flag.IntVar(&Cntxt.ListenPort, "listenport", listenPortDefault, listenPortUsage)
	flag.IntVar(&Cntxt.ListenPort, "l", listenPortDefault, listenPortUsage+" (shorthand)")
	flag.StringVar(&Cntxt.APIHost, "apiHost", apiHostDefault, apiHostUsage)
	flag.StringVar(&Cntxt.APIHost, "h", apiHostDefault, apiHostUsage+" (shorthand)")
	flag.StringVar(&Cntxt.ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&Cntxt.ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
}

// Process application (command line) flags. Note this happens automatically.
// No need to explicitly call this function (in fact that is a bad idea).
func init() {
	initFlags()
	flag.Parse()
	log.Printf("Initialized app package.")
}
