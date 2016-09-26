// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package middleware

import (
	"log"
	"net/http"
)

// LogrequestHandler will log the HTTP requests.
func LogrequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.URL)
		next.ServeHTTP(w, r)
	})
}
