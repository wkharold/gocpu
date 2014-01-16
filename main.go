// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	mode       = flag.String("mode", "server", "server or client")
	listenaddr = flag.String("listenaddr", ":1234", "Server port")
	addr       = flag.String("addr", "127.0.0.1:5640", "network address")
	debug      = flag.Int("d", 0, "print debug messages")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocpu ")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	switch *mode {
	case "server":
		server()
	case "client":
		client()
	case "runner":
		runner()
	default:
		usage()
	}

}
