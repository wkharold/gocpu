package main

import (
	"crypto/tls"
)

func keys(certfile, keyfile string) (cert tls.Certificate, err error) {
	cert, err = tls.LoadX509KeyPair(certfile, keyfile)
	return
}
