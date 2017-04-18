package rcache

import (
	"fmt"
	"io"
)

func writeBulkString(w io.Writer, b []byte) error {
	var err error
	if b == nil {
		_, err = w.Write([]byte("$-1\r\n"))
		return err
	}

	if _, err = w.Write([]byte(fmt.Sprintf("$%d\r\n", len(b)))); err != nil {
		return err
	}
	if _, err = w.Write(b); err != nil {
		return err
	}
	_, err = w.Write(crlf)
	return err
}
