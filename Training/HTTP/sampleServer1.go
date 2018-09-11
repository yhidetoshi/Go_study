package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, s)
}

func main(){
	http.Handle("/", String("Hello world"))
	http.ListenAndServe("localhost:8000", nil)
}

/*
$ curl localhost:8000/                                         ⏎
Hello world
*/