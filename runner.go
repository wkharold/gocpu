package main

/*
#include <stdlib.h>
#include <pty.h>
#include <utmp.h>
#cgo LDFLAGS: -lutil
*/
import "C"

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func runner() {
	var err error
	var am, as C.int

	if unshare() < 0 {
		log.Fatal("unshare failed")
	}

	res := C.openpty(&am, &as, nil, nil, nil)
	if res < 0 {
		log.Fatalf("openpty %v\n", res)
	}

	amaster := os.NewFile(uintptr(am), "/dev/ptmx")
	aslave := os.NewFile(uintptr(as), "slave")

	netreader := bufio.NewReader(os.Stdin)
	line, _, err := netreader.ReadLine()

	info := strings.Split(string(line), ":")
	if len(info) != 2 {
		log.Fatalf("Bad mount info '%v'\n", line)
	}
	err = syscall.Mount(info[0], "/mnt/term/", "9p", 0xc0ed0000, "port="+info[1])
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("/mnt/term/bin/bash")
	cmd.Stdin = aslave
	cmd.Stdout = aslave
	cmd.Stderr = aslave
	go io.Copy(amaster, os.Stdin)
	go io.Copy(os.Stdout, amaster)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setctty: true,
		Setsid:  true,
	}
	e := cmd.Run()
	if e != nil {
		log.Print(e)
	}
	return
}
