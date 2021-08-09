package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zli0111/video_server/api/session"
	"net/http"
)

// build a middleware to handle the auth of the user
type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

// the flow of our server
// main -> middleware -> defs(message, err) -> handlers -> dbops -> response
func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
