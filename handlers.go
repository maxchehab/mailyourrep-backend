package main

import (
	"fmt"
	"net/http"
)

//HelloWorld is the default handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
