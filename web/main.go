package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)

	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)

	router.POST("/userhome", userHomeHandler)

	router.POST("/api", apiHandler)

	router.GET("/videos/:vid-id", proxyHandler)

	router.POST("/upload/:vid-id", proxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}

// using proxy to transfer http://127.0.0.1:8080/upload/:vid-id
// into                    http://127.0.0.1:9000/upload/:vid-id
// to avoid cross origin
