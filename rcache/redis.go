package rcache

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

const debug = false

var (
	crlf = []byte("\r\n")
	lf   = byte('\n')
)

type Server struct {
	cache Cache
}

func NewServer(cache Cache) *Server {
	return &Server{
		cache: cache,
	}
}

func debugf(format string, args ...interface{}) {
	if debug {
		log.Printf(format, args...)
	}
}

func (server *Server) HandleConnection(c io.ReadWriter) error {
	r := bufio.NewReader(c)
	w := c // bufio.NewWriter(c)

	for {
		p, err := r.ReadByte()
		if err != nil {
			return err
		}

		switch p {
		case '*':
			a, err := readArray(r)
			if err != nil {
				return err
			}

			switch string(a[0].([]byte)) {
			case "SET":
				key := string(a[1].([]byte))
				value := a[2].([]byte)
				debugf("SET %s %s", key, value)
				server.cache.Set(key, value)
				if _, err = w.Write([]byte("+OK\r\n")); err != nil {
					return err
				}

			case "GET":
				key := string(a[1].([]byte))
				value := server.cache.Get(key)
				if value == nil {
					debugf("GET %s -> nil", key)
					writeBulkString(w, nil)
				} else {
					debugf("GET %s -> %s", key, value)
					writeBulkString(w, value.([]byte))
				}

			default:
				return fmt.Errorf("unexpected command %s", a[0])
			}

		default:
			return fmt.Errorf("unexpected command prefix %q", p)
		}

		// if err = w.Flush(); err != nil {
		// 	return err
		// }
	}
}
