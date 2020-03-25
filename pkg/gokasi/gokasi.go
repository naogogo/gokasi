package gokasi

/*
#cgo CFLAGS: -std=c99 -Wall
#cgo LDFLAGS: -lkakasi
#include <libkakasi.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import (
	"errors"
)

func Init(argv ...string) error {
	if len(argv) == 0 {
		// Arguments to Kakasi; see 'man kakasi'.
		argv = []string{
			"kakasi",
			"-Ha",
			"-Ka",
			"-Ja",
			"-Ea",
			"-ka",
			"-s",
			"-isjis",
			"-osjis",
		}
	}
	if err := kakasi_getopt_argv(argv); err != nil {
		return err
	}

	return nil
}

func New(s string) (string, error) {
	return kakasi_do_sjis(s)
}

func Destroy() error {
	return kakasi_close_kanwadict()
}

/* Package internal functions that interact with libkakasi below */

func kakasi_getopt_argv(argv []string) error {
	argc := len(argv)
	var argvcs [](*C.char)

	for _, i := range argv {
		argvcs = append(argvcs, C.CString(i))
	}

	rc := C.kakasi_getopt_argv(C.int(argc), &argvcs[0])
	for _, i := range(argvcs) {
		C.free(unsafe.Pointer(i))
	}

	if rc != 0 {
		return errors.New("kakasi_getopt_argv")
	}

	return nil
}

func kakasi_do(s string) (string, error) {
	cs := C.CString(s)
	ascii := C.kakasi_do(cs)

	C.free(unsafe.Pointer(cs))

	if ascii == nil {
		return "", errors.New("kakasi_do")
	}

	goascii := C.GoString(ascii)
	kakasi_free(ascii)

	return goascii, nil
}

// Same as kakasi_do, but encode to SJIS before conversion, and decode back
// afterwards. Kakasi does not seem to work properly with UTF-8 encoding.
func kakasi_do_sjis(s string) (string, error) {
	sjis, err := SJISEncode(s)
	if err != nil {
		return "", err
	}

	ascii, err := kakasi_do(sjis)
	if err != nil {
		return "", err
	}

	return SJISDecode(ascii)
}

func kakasi_close_kanwadict() error {
	if C.kakasi_close_kanwadict() != (C.int)(0) {
		return errors.New("kakasi_close_kanwadict")
	}
	return nil
}

// The C function returns 1 if an object was freed and 0 if it was a no-op, so
// use this as a void function also
func kakasi_free(s *C.char) {
	C.kakasi_free(s)
}
