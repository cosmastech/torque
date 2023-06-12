package fullstory

import (
	_ "github.com/tylermmorton/tmpl"
)

//go:generate tmpl bind ./fullstory.go --mode=embed

// Snippet is the dot context of the fullstory template.
// This template contains the fullstory javascript snippet.
//
//tmpl:bind fullstory.tmpl.html
type Snippet struct {
	OrgId string
}
