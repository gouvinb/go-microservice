// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"log"
	"net/http"
	"route/routewrapper"

	"github.com/josephspurrier/csrfbanana"

	"route/middleware"
	"shared"
	"time"
)

func init() {
	routewrapper.Get("/sample/acl", routewrapper.Chain(SampleACLGET))
	routewrapper.Post("/sample/acl", routewrapper.Chain(SampleACLPOST))
	routewrapper.Get("/sample/acl/auth", routewrapper.Chain(SampleACLAuthGET, middleware.DisallowAnon))
	routewrapper.Get("/sample/acl/anon", routewrapper.Chain(SampleACLAnonGET, middleware.DisallowAuth))
}

// SampleACLGET displays the default acl page.
func SampleACLGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "sample/acl"
		v.Vars["token"] = csrfbanana.Token(w, r, sess)
		if sess.Values["id"] == nil {
			v.Vars["id"] = "id nil"
		} else {
			v.Vars["id"] = sess.Values["id"]
		}
		// Refill any form fields from a POST operation
		shared.ViewRepopulate([]string{"session"}, r.Form, v.Vars)
		v.ViewRender(w)
	} else {
		msg := `{"state":"ok","timestamp":"` + time.Now().String() + `","path":"` +
			r.URL.Path + `"}`
		fmt.Fprint(w, msg)
	}
}

// SampleACLPOST displays the default acl page.
func SampleACLPOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		sessionValue := r.FormValue("session")
		shared.SessionEmpty(sess)
		if sessionValue == "true" {
			sess.Values["id"] = "1337" // or id of user
			log.Println(sess.Save(r, w).Error())
		} else if sessionValue == "false" {
			sess.Save(r, w)
		}
		http.Redirect(w, r, "/sample/acl", http.StatusFound)
		return
	}
}

// SampleACLAuthGET displays the default acl page.
func SampleACLAuthGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "sample/acl/auth"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		msg := `{"state":"ok","timestamp":"` + time.Now().String() + `","path":"` +
			r.URL.Path + `"}`
		fmt.Fprint(w, msg)
	}
}

// SampleACLAnonGET displays the default acl page.
func SampleACLAnonGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "sample/acl/anon"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		msg := `{"state":"ok","timestamp":"` + time.Now().String() + `","path":"` +
			r.URL.Path + `"}`
		fmt.Fprint(w, msg)
	}
}
