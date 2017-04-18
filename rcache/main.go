// +build ignore

package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	. "."
)

func runHTTPHandler() {
	const addr = "127.0.0.1:8080"
	log.Printf("HTTP: Listening on %s ...", addr)
	log.Printf("\thttp://%s/debug/pprof", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	const addr = "127.0.0.1:6379"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TCP: Listening on %s ...", addr)
	defer l.Close()

	go runHTTPHandler()

	server := NewServer(NewMap(0))

	for {
		c, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go func() {
			err := server.HandleConnection(c)
			if err != nil {
				log.Print(err)
			}
			c.Close()
		}()
	}
}
