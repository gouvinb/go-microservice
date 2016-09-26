// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"log"
	"net/http"

	"route/middleware"
	"route/router"
	"shared"
)

func init() {
	log.Println("Init index handlers")

	router.Get("/api/test", router.Chain(TestAPIGET))
	router.Get("/api/test/auth", router.Chain(TestAPIAuthGET, middleware.DisallowAnon))
	router.Get("/api/test/anon", router.Chain(TestAPIAnonGET, middleware.DisallowAuth))
	router.Post("/test/anon", router.Chain(TestAPIAnonGET, middleware.DisallowAuth))
}

// TestAPIGET displays the default home page.
func TestAPIGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"your are in GET /api/test and all user are allowed\"}"
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}

// TestAPIAuthGET displays the default home page.
func TestAPIAuthGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"your are in GET /api/test and authenticated user are allowed only\"}"
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}

// TestAPIAnonGET displays the default home page.
func TestAPIAnonGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"your are in GET /api/test and anonymous user are allowed only\"}"
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}
