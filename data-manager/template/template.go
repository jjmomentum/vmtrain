// Package template holds templates for all the packages.
//
// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package template

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vmtrain/data-manager/models"
	"github.com/vmtrain/data-manager/stats"
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

// Retrieve all Message entries from a topic via REST call.
func (t *Template) getTopic(topic string) []models.Message {
	url := t.generateAPIUrl("/api/topic/" + topic + "?peekall=true")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	var msgs []models.Message
	if err := json.Unmarshal(body, &msgs); err != nil {
		log.Println(err)
	}

	return msgs
}

// Retrieve all Topic names via REST call.
func (t *Template) getAllTopics() *models.Topics {
	url := t.generateAPIUrl("/api/topics")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	topics := &models.Topics{}
	topics.FromJSON(body)

	return topics
}

// Add a topic message, to the models via REST call.
func (t *Template) postTopic(topic string, msg models.Message) {
	jsonData, err := json.Marshal(msg)
	perror(err)

	url := t.generateAPIUrl("/api/topic/" + topic)
	log.Println("url: " + url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	perror(err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	_, err = ioutil.ReadAll(rsp.Body)
	perror(err)
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
