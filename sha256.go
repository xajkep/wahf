/*

COMPILATION
===========
GOOS=js GOARCH=wasm go build -o static/sha256.wasm sha256.go

wasm_exec.js
============
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static

*/
package main

import (
	"crypto"
	_ "crypto/sha256"
	"encoding/hex"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("SHA256", js.FuncOf(hash))
	<-done
}

func hash(this js.Value, args []js.Value) interface{} {
	h := crypto.SHA256.New()
	h.Write([]byte(args[0].String()))
	return hex.EncodeToString(h.Sum(nil))
}
