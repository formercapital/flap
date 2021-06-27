// +build !detached

package main

import (
	"strconv"

	"github.com/errtokyo/flap/go/server"
)

import "C"

//export serve
func serve(port C.int, dir *C.char) {
	p := strconv.Itoa(int(port))
	d := C.GoString(dir)
	server.Serve(p, d)
}

func main() {}
