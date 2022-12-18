// Package basics contains some basic knowledge.
package basics

import "errors"

// LearnGoDoc is the main function of this package.
func LearnGoDoc() {

	// call " godoc -http=:6060" from command line
	// open a web browser and open "http://localhost:6060" and look for your program
}

// Below are just some elements (struct and variables) to test the documentation of them.

// Server contains basic information about servers.
type Server struct {
	port string
	ip   string
}

// ErrConnectingServer error when server cant connect
// ErrRequestTimeout error when server gets a timeout
var (
	ErrConnectingServer = errors.New("can't connect to server")
	ErrRequestTimeout   = errors.New("request took to long")
)
