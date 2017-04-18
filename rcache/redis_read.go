package rcache

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

func readCRLF(r *bufio.Reader) error {
	b := make([]byte, 2)
	if _, err := io.ReadFull(r, b); err != nil {
		return err
	}
	if !reflect.DeepEqual(b, []byte("\r\n")) {
		return errors.New("crlf not found")
	}
	return nil
}

func readSimpleString(r *bufio.Reader) (string, error) {
	b, err := r.ReadBytes(lf)
	if err != nil {
		return "", err
	}
	if !bytes.HasSuffix(b, crlf) {
		return "", errors.New("crlf not found")
	}
	return string(b[:len(b)-2]), nil
}

func readInteger(r *bufio.Reader) (int64, error) {
	s, err := readSimpleString(r)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(s, 10, 64)
}

func readBulkString(r *bufio.Reader) ([]byte, error) {
	n, err := readInteger(r)
	if err != nil {
		return nil, err
	}

	b := make([]byte, n)
	if _, err = io.ReadFull(r, b); err != nil {
		return nil, err
	}
	if err = readCRLF(r); err != nil {
		return nil, err
	}
	return b, nil
}

func readArray(r *bufio.Reader) ([]interface{}, error) {
	n, err := readInteger(r)
	if err != nil {
		return nil, err
	}

	res := make([]interface{}, n)
	for i := int64(0); i < n; i++ {
		p, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		switch p {
		case '$':
			b, err := readBulkString(r)
			if err != nil {
				return nil, err
			}
			res[i] = b

		default:
			return nil, fmt.Errorf("unexpected array element prefix %q", p)
		}
	}
	return res, nil
}
