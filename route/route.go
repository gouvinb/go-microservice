// Package route load and list all route for web server
package route

import (
	"log"
	"net/http"

	"github.com/gouvinb/go-microservice/controller"
	"github.com/gouvinb/go-microservice/route/middleware"
	"github.com/gouvinb/go-microservice/shared"

	"github.com/gorilla/context"
	"github.com/josephspurrier/csrfbanana"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load returns the routes and middleware
func Load() http.Handler {
	log.Println("Load all handlers")
	return middlewareHandler(routes())
}

// LoadHTTPS returns the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	log.Println("Load HTTPS handlers")
	return middlewareHandler(routes())
}

// LoadHTTP returns the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	log.Println("Load HTTP handlers")
	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
	return middlewareHandler(routes())
}

// routes list all routes
func routes() *httprouter.Router {
	r := httprouter.New()

	log.Println("Set 404 handler")
	r.NotFound = alice.
		New().
		ThenFunc(controller.Error404)

	log.Println("Set index handler")
	r.GET("/", middleware.Handler(alice.
		New().
		ThenFunc(controller.Index)))

	log.Println("Set Pprof handler")
	r.GET("/debug/pprof/*pprof", middleware.Handler(alice.
		New(middleware.DisallowAnon).
		ThenFunc(middleware.PprofHandler)))

	return r
}

// middlewareHandler for prevents CSRF and Double Submits
func middlewareHandler(h http.Handler) http.Handler {
	log.Println("Prevents CSRF and Double Submits")
	cs := csrfbanana.New(h, shared.Store, shared.Name)
	cs.FailureHandler(http.HandlerFunc(controller.InvalidToken))
	cs.ClearAfterUsage(true)
	cs.ExcludeRegexPaths([]string{"/static(.*)"})
	csrfbanana.TokenLength = 32
	csrfbanana.TokenName = "token"
	csrfbanana.SingleToken = false
	h = cs

	log.Println("Logger request activated")
	h = middleware.LogrequestHandler(h)

	log.Println("Clear handler for Gorilla Context")
	h = context.ClearHandler(h)

	return h
}
