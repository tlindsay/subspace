package main

import (
	"fmt"

	"github.com/fastly/compute-sdk-go/fsthttp"
	adapter "github.com/joeshaw/fsthttp-adapter/handler"
	subspace "github.com/tlindsay/subspace/api"
)

// The entry point for your application.
//
// Use this function to define your main request handling logic. It could be
// used to route based on the request properties (such as method or path), send
// the request to a backend, make completely new requests, and/or generate
// synthetic responses.

func Error(w fsthttp.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

func main() {
	h := subspace.Handler()
	adapted := adapter.Adapt(h)
	fsthttp.Serve(adapted)
}
