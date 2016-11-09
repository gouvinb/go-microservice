// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"log"
	"net/http"
	"route/routewrapper"

	"shared"
	"time"
)

func init() {
	log.Println("Init index handlers")

	routewrapper.Get("/", routewrapper.Chain(Index))
}

// Index displays the default home page.
func Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.SessionInstance(r)

	if sess != nil {
		v := shared.ViewNew(r)
		v.Name = "index/index"
		v.ViewRender(w)
		sess.Save(r, w)
	} else {
		msg := `{"state":"ok","timestamp":"` + time.Now().String() + `","path":"` +
			r.URL.Path + `"}`
		fmt.Fprint(w, msg)
	}
}
