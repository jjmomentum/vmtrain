// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import "net/http"

// JobCreate creates a job and stores it in etcd
// Request URL:
// 		/jobs
// Request body:
// {
//  apiVersion: <v>
//  kind: GeronimoJob
//  spec: {
//    system: <system>
//    environment: <environ>
//  }
// }
// Response:
// {
//  apiVersion: <v>
//  kind: GeronimoJob
//  metadata: {
//    uuid: <id>
//    selfLink: /job/<id>
//    creationTimestamp: <iso8601>
//    clusterAddress: <ip address>
//  }
//  spec: {
//    system: <system>
//    environment: <environment>
//  }
//  status: {
//    phase: PENDING|ACTIVE|COMPLETED|FAILED
//    startTimestamp: <iso8601>
//    endTimestamp: <iso8601>
//  }
// }
func CreateServer(w http.ResponseWriter, r *http.Request) {
	//TODO add logic to create a server and store it in the blob service
}

// JobCreate creates a job and stores it in etcd
// Request URL:
// 		/jobs
// Request body:
// {
//  apiVersion: <v>
//  kind: GeronimoJob
//  spec: {
//    system: <system>
//    environment: <environ>
//  }
// }
// Response:
// {
//  apiVersion: <v>
//  kind: GeronimoJob
//  metadata: {
//    uuid: <id>
//    selfLink: /job/<id>
//    creationTimestamp: <iso8601>
//    clusterAddress: <ip address>
//  }
//  spec: {
//    system: <system>
//    environment: <environment>
//  }
//  status: {
//    phase: PENDING|ACTIVE|COMPLETED|FAILED
//    startTimestamp: <iso8601>
//    endTimestamp: <iso8601>
//  }
// }
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	// TODO add logic to create a reservation and add it to the blob service
}

// JobsListShow displays a list of jobs stored in etcd
// Request URL:
// 		/jobs
// Response:
// {
//  apiVersion: <v>
//  kind: GeronimoJobsList
//  metadata: {}
//  items: [
//    <GeronimoJob>
//    <GeronimoJob>
//    <GeronimoJob>
//  ]
// }
func ShowReservationList(w http.ResponseWriter, r *http.Request) {
	// TODO add logic to show a list of reservations
}
