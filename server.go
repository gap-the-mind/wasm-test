// +build !js

package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func build() {
	log.Print("Building wasm")
	cmd := exec.Command("go", "build", "-o", "www/main.wasm")

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOOS=js")
	cmd.Env = append(cmd.Env, "GOARCH=wasm")

	cmd.Run()
}

func main() {
	build()

	log.Print("Starting server")
	http.ListenAndServe(":8080", http.FileServer(http.Dir("www")))
}
