// +build js

package main

import (
	"log"
	"syscall/js"

	"github.com/go-git/go-billy/v5/osfs"
)

func main() {
	ch := make(chan struct{})

	fs := osfs.New("/")
	_, err := fs.Create("test.txt")

	log.Print(err)

	<-ch
	js.Global().Get("console").Call("log", "Exiting Wasm module ...")
}
