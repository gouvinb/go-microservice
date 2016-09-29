// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package controller

import (
	"fmt"
	"log"
	"net/http"

	"route/router"
	"shared"
)

func init() {
	log.Println("Init index handlers")

	router.Get("/", router.Chain(Index))
}

// Index displays the default home page.
func Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := `{"message": "if you see this json, it's because the micro service is OP"}`
	fmt.Fprint(w, msg)

	if sess != nil {
		sess.Save(r, w)
	}
}
