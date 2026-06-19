package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct{ 
	Name string `json:"name"`
}

var userCache = make(map[int] User)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/",handleRoot)
	mux.HandleFunc("POST /users",createUser)


	fmt.Println("Server listening on 8080")

	http.ListenAndServe(":8000",mux)
}


func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello")

}


func createUser(w http.ResponseWriter, r *http.Request){
	var user User
	err:= json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return 
	}

	if user.Name !=""{

	}
}