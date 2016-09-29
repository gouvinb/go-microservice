// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package middleware

import (
	"log"
	"net/http"

	"shared"
)

// DisallowAuth does not allow authenticated users to access the page.
func DisallowAuth(h http.Handler) http.Handler {
	log.Println("Set DisallowAuth")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session
		if sess := shared.SessionInstance(r); sess != nil {
			// If user is authenticated, don't allow them to access the page
			if sess.Values["id"] != nil {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// DisallowAnon does not allow anonymous users to access the page.
func DisallowAnon(h http.Handler) http.Handler {
	log.Println("Set DisallowAnon")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session
		if sess := shared.SessionInstance(r); sess != nil {
			// If user is not authenticated, don't allow them to access the page
			if sess.Values["id"] == nil {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
