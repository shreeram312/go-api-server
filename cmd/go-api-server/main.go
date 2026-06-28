package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
)

type api struct{
	addr string
}


type User struct{ 
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var users = []User{}

func (a * api) getUserHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")


	err:=json.NewEncoder(w).Encode(users)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusOK)
	// w.Write([] byte("hello"))
	// fmt.Fprintf(w,"hello br")
	// both are allowed 
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


