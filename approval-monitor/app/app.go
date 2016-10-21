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
		dataManagerUsage               = "base URL endpoint for the data-manager microservice."
		dataManagerDefault             = "http://localhost:6001"
		approvalPollIntervalSecUsage   = "interval of time in between polls to the approvals service in seconds"
		approvalPollIntervalSecDefault = 60
	)

	flag.StringVar(&Cntxt.DataManagerURL, "data-manager", dataManagerDefault, dataManagerUsage+" (shorthand)")
	flag.IntVar(&Cntxt.ApprovalPollIntervalSec, "poll-interval", approvalPollIntervalSecDefault, approvalPollIntervalSecUsage+" (shorthand")
}

// Process application (command line) flags. Note this happens automatically.
// No need to explicitly call this function (in fact that is a bad idea).
func init() {
	initFlags()
	flag.Parse()
	log.Printf("Initialized app package.")
}
