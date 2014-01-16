package main

/*
#include <stdlib.h>
#include <pty.h>
#include <utmp.h>
#include <curses.h>
#cgo LDFLAGS: -lutil -lcurses
*/
import "C"

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func client() {
	go runufs()
	addrstring := fmt.Sprintf("%v\n", *addr)
	fmt.Printf("serving 9p on %v\n", *addr)
	c, err := net.Dial("tcp", *listenaddr)
	if err != nil {
		log.Fatal(err)
	}
	c.Write([]byte(addrstring))

	C.initscr()
	C.raw()

	go io.Copy(c, os.Stdout)
	io.Copy(os.Stdin, c)
	C.endwin()
}
