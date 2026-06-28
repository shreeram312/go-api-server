package main

import (
	"fmt"
	"log"
	"net/http"
)

type api struct{
	addr string
}


func (a * api) getUserHandler(w http.ResponseWriter, r *http.Request){
	// w.Write([] byte("hello"))
	fmt.Fprintf(w,"hello br")
}


func main(){
	api:=&api{addr: ":8080"}
	mux:= http.NewServeMux()


	srv:=&http.Server{
		Addr: api.addr,
		Handler: mux,
	}
	mux.HandleFunc("GET /users", api.getUserHandler)

	err:= srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

	
}


