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

	router.Get("/api/sample", router.Chain(SampleAPIGET))
	router.Get("/api/sample/auth", router.Chain(SampleAPIAuthGET, middleware.DisallowAnon))
	router.Get("/api/sample/anon", router.Chain(SampleAPIAnonGET, middleware.DisallowAuth))
	if shared.Name != "" && shared.Store != nil {
		router.Post("/sample/anon", router.Chain(SampleAPIAnonGET, middleware.DisallowAuth))
	}
}

// SampleAPIGET displays the default home page.
func SampleAPIGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"your are in GET /api/sample and all user are allowed\"}"
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}

// SampleAPIAuthGET displays the default home page.
func SampleAPIAuthGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"your are in GET /api/sample and authenticated user are allowed only\"}"
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}

// SampleAPIAnonGET displays the default home page.
func SampleAPIAnonGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"your are in GET /api/sample and anonymous user are allowed only\"}"
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}
