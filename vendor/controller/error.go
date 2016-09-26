// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"log"
	"net/http"

	"route/router"
)

func init() {
	log.Println("Init error handlers")

	// This does not work for routes where the path matches, but the method does not
	// (on HEAD and OPTIONS need to check)
	// https://github.com/julienschmidt/httprouter/issues/13
	var e405 http.HandlerFunc = Error405
	router.Instance().HandleMethodNotAllowed = true
	router.Instance().MethodNotAllowed = e405

	// 404 Page
	var e404 http.HandlerFunc = Error404
	router.Instance().NotFound = e404
}

// Error404 handles 404 - Page Not Found.
func Error404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "{ \"error\": \"Not Found 404\"}")
}

// Error405 handles 405 - Page Not Found.
func Error405(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "{ \"error\": \"Method Not Allowed 405\"}")
}

// Error500 handles 500 - Internal Server Error.
func Error500(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "{ \"error\": \"Internal Server Error 500\"}")
}

// InvalidToken handles CSRF attacks.
func InvalidToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, "{ \"error\": \"Token expired\"}")
}
