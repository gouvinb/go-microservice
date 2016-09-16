// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"net/http"

	"github.com/gouvinb/go-microservice/route/router"
	// "github.com/gouvinb/go-microservice/shared"
)

func init() {
	// This does not work for routes where the path matches, but the method does not
	// (on HEAD and OPTIONS need to check)
	// https://github.com/julienschmidt/httprouter/issues/13
	//var e405 http.HandlerFunc = Error405
	//router.Instance().HandleMethodNotAllowed = true
	//router.Instance().MethodNotAllowed = e405

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
