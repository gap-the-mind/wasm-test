// +build js

package main

import (
	"log"
	"syscall/js"

	"github.com/go-git/go-billy/v5/osfs"
)

func main() {
	ch := make(chan struct{})

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "message")
	document.Get("body").Call("appendChild", p)

	fs := osfs.New("/")
	_, err := fs.Create("/test.txt")

	log.Fatal(err)

	<-ch
	js.Global().Get("console").Call("log", "Exiting Wasm module ...")
}
