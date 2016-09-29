// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package shared

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// Store is the cookie store
	Store *sessions.CookieStore
	// Name is the session name
	Name string
)

// Session stores session level information.
type Session struct {
	// Pulled from: http://www.gorillatoolkit.org/pkg/sessions#Options
	Options sessions.Options `json:"Options"`
	// Name for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.Get
	Name string `json:"Name"`
	// Key for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.New
	SecretKey string `json:"SecretKey"`
	// Bool for session enabler
	EnableSession bool `json:"EnableSession"`
}

// SessionConfigure the session cookie store.
func SessionConfigure(s Session) {
	if IsSessionEnabled(s) {
		Store = sessions.NewCookieStore([]byte(GetSessionSecretKey(s)))
		Name = GetSessionName(s)

		s.Options.Path = GetSessionOptionPath(s)
		s.Options.Domain = GetSessionOptionDomain(s)
		s.Options.MaxAge = GetSessionOptionMaxAge(s)
		s.Options.Secure = GetSessionOptionSecure(s)
		s.Options.HttpOnly = GetSessionOptionHTTPOnly(s)

		Store.Options = &s.Options
	}
}

// Instance returns a new session, never returns an error.
func SessionInstance(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, Name)
	return session
}

// SessionEmpty deletes all the current session values.
func SessionEmpty(sess *sessions.Session) {
	// Clear out all stored values in the cookie
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}
