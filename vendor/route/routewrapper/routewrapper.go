// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package routewrapper

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"
)

var (
	r RouteWapperInfo
)

const (
	params = "params"
)

// RouteWapperInfo is the details.
type RouteWapperInfo struct {
	Router *httprouter.Router
}

func init() {
	// Init router
	r.Router = httprouter.New()
}

// ReadConfig returns the information.
func ReadConfig() RouteWapperInfo {
	log.Println("Read config")
	return r
}

// Instance returns the router.
func Instance() *httprouter.Router {
	return r.Router
}

// Params returns the URL parameters.
func Params(r *http.Request) httprouter.Params {
	log.Println("Get params from router")
	return context.Get(r, params).(httprouter.Params)
}

// Chain returns handle with chaining using Alice.
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	return Handler(alice.New(c...).ThenFunc(fn))
}
