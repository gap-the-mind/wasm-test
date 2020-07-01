// +build !js

package main

import (
	"os"
	"os/exec"
)

func build() {
	cmd := exec.Command("go", "build", "-o", "main.wasm")

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOOS=js")
	cmd.Env = append(cmd.Env, "GOARCH=wasm")

	cmd.Run()
}

func main() {
	build()
	// http.ListenAndServe(":8080", http.FileServer(http.Dir(`.`)))
}
