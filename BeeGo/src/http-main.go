package main

import (
	"io"
	"log"
	"net/http"
)


func main(){
	//set routing
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal(err)
	}
}


func sayHello(w http.ResponseWriter,r *http.Request){
io.WriteString(w,"Hello world, this is version 1。")
}