package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	//TODO: Add users and todos handlers
	//TODO: Decide where to create the users and todos handlers

	return mux
}
