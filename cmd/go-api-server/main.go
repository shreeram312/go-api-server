package main

import (
	"log"
	"net/http"
)


type server struct {
	addr string
}


func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello "))	
}


func main(){
	s:=&server{addr: ":8080"}
	err:=http.ListenAndServe(s.addr,s)
	if err !=nil{
		log.Fatal(err)
	}

}