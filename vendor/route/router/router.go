package router

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"
)

var (
	r RouterInfo
)

const (
	params = "params"
)

// RouterInfo is the details.
type RouterInfo struct {
	Router *httprouter.Router
}

func init() {
	log.Println("Init router")
	r.Router = httprouter.New()
}

// ReadConfig returns the information.
func ReadConfig() RouterInfo {
	log.Println("Read config")
	return r
}

// Instance returns the router.
func Instance() *httprouter.Router {
	log.Println("Instance router called")
	return r.Router
}

// Params returns the URL parameters.
func Params(r *http.Request) httprouter.Params {
	log.Println("Get params from router")
	return context.Get(r, params).(httprouter.Params)
}

// Chain returns handle with chaining using Alice.
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	log.Println("Chain handler with alice")
	return Handler(alice.New(c...).ThenFunc(fn))
}
