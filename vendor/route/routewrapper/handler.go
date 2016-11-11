package routewrapper

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// HandlerFunc accepts the name of a function so you don't have to wrap it with
// http.HandlerFunc.
// Example: routerwrapper.Get("/", routerwrapper.HandlerFunc(controller.Index))
// Source: http://nicolasmerouze.com/guide-routers-golang/
func HandlerFunc(h http.HandlerFunc) httprouter.Handle {
	log.Println("HandlerFunc used")
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, params, p)
		h.ServeHTTP(w, r)
	}
}

// Handler accepts a handler to make it compatible with http.HandlerFunc.
// Example: routerwrapper.Get("/", routerwrapper.Handler(http.HandlerFunc(
// controller.Index)))
// Source: http://nicolasmerouze.com/guide-routers-golang/
func Handler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, params, p)
		h.ServeHTTP(w, r)
	}
}
