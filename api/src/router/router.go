package router

import "github.com/gorilla/mux"

func Generator() *mux.Router {
	return mux.NewRouter()
}