package gokasi

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
)

// https://gist.github.com/hyamamoto/db03c03fd624881d4b84

func SJISDecode(s string) (string, error) {
	return readAll(strings.NewReader(s), japanese.ShiftJIS.NewDecoder())
}

func SJISEncode(s string) (string, error) {
	return readAll(strings.NewReader(s), japanese.ShiftJIS.NewEncoder())
}

func readAll(reader io.Reader, transformer transform.Transformer) (string, error) {
	s, err := ioutil.ReadAll(transform.NewReader(reader, transformer))
	if err == nil {
		return string(s), nil
	}
	return "", err
}
