// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// Package route load router for web server
package route

import (
	"log"
	"net/http"

	"controller"
	"route/middleware"
	"route/router"
	"shared"

	"github.com/gorilla/context"
	"github.com/josephspurrier/csrfbanana"
)

// Load returns the routes and middleware.
func Load() http.Handler {
	log.Println("Load all handlers")
	return middlewareHandler(router.Instance())
}

// LoadHTTPS returns the HTTP routes and middleware.
func LoadHTTPS() http.Handler {
	log.Println("Load HTTPS handlers")
	return middlewareHandler(router.Instance())
}

// LoadHTTP returns the HTTPS routes and middleware.
func LoadHTTP() http.Handler {
	log.Println("Load HTTP handlers")
	// Uncomment this and comment out the line above to always redirect to HTTPS
	// return http.HandlerFunc(redirectToHTTPS())
	return middlewareHandler(router.Instance())
}

// redirectToHTTPS redirect from HTTP to HTTPS.
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	log.Println("Redirect to https")
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// middlewareHandler for prevents CSRF and Double Submits.
func middlewareHandler(h http.Handler) http.Handler {
	log.Println("Prevents CSRF and Double Submits")
	log.Println("##", shared.Name, "##")
	log.Println("##", shared.Store, "##")
	if shared.Name != "" && shared.Store != nil {
		cs := csrfbanana.New(h, shared.Store, shared.Name)
		cs.FailureHandler(http.HandlerFunc(controller.InvalidToken))
		cs.ClearAfterUsage(true)
		cs.ExcludeRegexPaths([]string{"/static(.*)", "/api(.*)"})
		csrfbanana.TokenLength = 32
		csrfbanana.TokenName = "token"
		csrfbanana.SingleToken = false
		h = cs
	}

	if shared.EnableCors {
		h = shared.Handler(h)
	}

	log.Println("Logger request activated")
	h = middleware.LogrequestHandler(h)

	log.Println("Clear handler for Gorilla Context")
	h = context.ClearHandler(h)

	return h
}
