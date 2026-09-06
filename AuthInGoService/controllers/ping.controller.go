package controllers

import "net/http"

func Pinghandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Pong"));
}