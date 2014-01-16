/*
 * gproc, a Go reimplementation of the LANL version of bproc and the LANL XCPU software.
 *
 * This software is released under the GNU Lesser General Public License, version 2, incorporated herein by reference.
 *
 * Copyright (2010) Sandia Corporation. Under the terms of Contract DE-AC04-94AL85000 with Sandia Corporation,
 * the U.S. Government retains certain rights in this software.
 */

package main

import (
	"syscall"
)

const (
	linuxhack = 0xc0ed0000
)

func unshare() int {
	_, _, syscallerr := syscall.Syscall(syscall.SYS_UNSHARE, uintptr(0x00020000), uintptr(0), uintptr(0))
	return int(syscallerr)
}
