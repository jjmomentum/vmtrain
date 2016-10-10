// Package template holds templates for all the packages.
//
// Copyright (c) 2015 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package template

import (
	"testing"
)

func TestGenerateAPIUrl(t *testing.T) {
	tmpl := New("/", "www.website.com/")

	expected := "http://www.website.com/testing"
	actual := tmpl.generateAPIUrl("testing")

	if expected != actual {
		t.Errorf("Test failed!, expected: '%s', got: '%s'", expected, actual)
	}
}
