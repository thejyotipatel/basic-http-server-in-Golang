package main

import (
	"basic-server/controller"
	"fmt"
	"log"
	"net/http"
)
  

func main() {
	
fileServer := http.FileServer(http.Dir("./public"))

http.Handle("/", fileServer)
 
http.HandleFunc("/home",  controller.HomeHandler)
http.HandleFunc("/form", controller.FormHandler)

fmt.Println("Starting server at port 8080")
if err := http.ListenAndServe(":8080", nil) ;err != nil {
	log.Fatal(err)
}    
}