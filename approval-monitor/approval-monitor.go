// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vmtrain/approval-monitor/app"
	"github.com/vmtrain/approval-monitor/models"
)

const (
	getApprovalsURL      = "http://approval.vmwaredevops.appspot.com/api/v1/approvables?approved=true&teamID=4"
	updateReservationURL = "api/reservations"
	deleteApprovalURL    = "http://approval.vmwaredevops.appspot.com/api/v1/approvables/"
	jsonHeader           = "application/json"
)

// Called by main, which is just a wrapper for this function. The reason
// is main can't directly pass back a return code to the OS.
func realMain() int {
	for {
		log.Printf("Querying for approvals...")
		var approvals models.ApprovalList
		_, err := app.MakeRequest(
			getApprovalsURL,
			http.MethodGet,
			jsonHeader,
			&approvals,
			nil,
			200,
		)
		if err != nil {
			log.Println("Shutdown caused by:" + err.Error())
			return 1
		}

		for _, approval := range approvals {
			if approval.Approved {
				var reservationResponse models.Reservation
				reservationPayload := models.Reservation{Approved: approval.Approved}
				// Unmarshal JSON
				b, err := json.Marshal(reservationPayload)
				if err != nil {
					log.Println("Shutdown caused by:" + err.Error())
					return 1
				}

				_, err = app.MakeRequest(
					fmt.Sprintf("%s/%s/%s", app.Cntxt.DataManagerURL, updateReservationURL, approval.Description),
					http.MethodPut,
					jsonHeader,
					&reservationResponse,
					bytes.NewReader(b),
					200,
				)
				if err != nil {
					log.Println("Shutdown caused by:" + err.Error())
					return 1
				}

				_, err = app.MakeRequest(
					fmt.Sprintf("%s%d", deleteApprovalURL, approval.ID),
					http.MethodDelete,
					jsonHeader,
					nil,
					nil,
					200,
					500,
				)
				if err != nil {
					log.Println("Shutdown caused by:" + err.Error())
					return 1
				}
			}
		}
		log.Printf("Sleeping for %d seconds\n", app.Cntxt.ApprovalPollIntervalSec)
		time.Sleep(time.Duration(uint(time.Second) * uint(app.Cntxt.ApprovalPollIntervalSec)))

	}
}

func main() {
	// Delegate to realMain so defered operations can happen (os.Exit exits
	// the program without servicing defer statements)
	os.Exit(realMain())
}
