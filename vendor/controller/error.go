// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"net/http"
	"route/routewrapper"
	"shared"
)

func init() {
	// This does not work for routes where the path matches, but the method does
	// not (on HEAD and OPTIONS need to check)
	// https://github.com/julienschmidt/httprouter/issues/13
	var e405 http.HandlerFunc = Error405
	routewrapper.Instance().HandleMethodNotAllowed = true
	routewrapper.Instance().MethodNotAllowed = e405

	// 404 Page
	var e404 http.HandlerFunc = Error404
	routewrapper.Instance().NotFound = e404
}

// Error404 handles 404 - Page Not Found.
func Error404(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "error/404"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{ \"error\": \"Not Found 404\"}")
	}
}

// Error405 handles 405 - Method Not Allowed.
func Error405(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "error/405"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "{ \"error\": \"Method Not Allowed 405\"}")
	}
}

// Error500 handles 500 - Internal Server Error.
func Error500(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "error/500"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{ \"error\": \"Internal Server Error 500\"}")
	}
}

// InvalidToken handles CSRF attacks.
func InvalidToken(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "error/csrf"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "{ \"error\": \"Token expired\"}")
	}
}
