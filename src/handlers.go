package main

import (
	"fmt"
	"net/http"
)

//HelloWorld is the default handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

//Send is the overlaying function that sends letters.
func Send(w http.ResponseWriter, r *http.Request) {
	requestData, err := CreateRequestData(r)
	if err != nil {
		Error(err)
		return
	}
	fmt.Fprintln(w, requestData)

	reps, err := GetReps(requestData)
	fmt.Fprintln(w, reps)
}
