package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// HandlerFunc accepts the name of a function so you don't have to wrap it with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.HandlerFunc(controller.Index))
func HandlerFunc(h http.HandlerFunc) httprouter.Handle {
	log.Println("Set handlerFunc")
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}

// Handler accepts a handler to make it compatible with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.Handler(http.HandlerFunc(controller.Index)))
func Handler(h http.Handler) httprouter.Handle {
	log.Println("Set handler")
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}