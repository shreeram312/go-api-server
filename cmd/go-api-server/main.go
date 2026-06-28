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


func (a *api) createUserHandler(w http.ResponseWriter, r * http.Request){
	var user User

err:=	json.NewDecoder(r.Body).Decode(&user)

	if err !=nil{
		http.Error(w,"Invalid json",http.StatusBadRequest)
	}	

	for _,existingUser := range users{
		if existingUser.FirstName == user.FirstName &&
		existingUser.LastName == user.LastName {
	
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	}  
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

}

func (a * api) getUserHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")


	
	

	w.WriteHeader(http.StatusOK)
	err:=json.NewEncoder(w).Encode(users)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

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
	mux.HandleFunc("POST /users", api.createUserHandler)

	err:= srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

	
}


