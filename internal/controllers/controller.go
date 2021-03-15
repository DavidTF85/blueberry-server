package controllers

import (
	"net/http"
	"fmt"
)

type Controller struct {

}

func New()(*Controller){
	return &Controller{}
}

func (c*Controller) HandleRequest (w http.ResponseWriter, r *http.Request) 