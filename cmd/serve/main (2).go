package main

import (
	//others
	"net/http"
	"log"
	"fmt"
    "time"

	// my folders
	"github.com\mulberry-server\internal\controllers"

)

func main (){
	c:controllers.New()

	//multiplexer
 	mux:= http.NewserveMux()
	 //conect to internet & my controller
	mux.HandleFunc("/", c.HandleRequest)
}