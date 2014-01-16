package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

func server() {
	l, err := net.Listen("tcp", *listenaddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Print(err)
		} else {
			go func() {
				cmd := exec.Command("gocpu", "-mode", "runner")
				cmd.Stdin = c
				cmd.Stdout = c
				cmd.Stderr = c
				err := cmd.Run()
				if err != nil {
					fmt.Print(err)
				}
				c.Close()
				return
			}()
		}
	}
}
