package router

import "github.com/julienschmidt/httprouter"

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func Delete(path string, fn httprouter.Handle) {
	r.Router.DELETE(path, fn)
}

// Get is a shortcut for router.Handle("GET", path, handle)
func Get(path string, fn httprouter.Handle) {
	r.Router.GET(path, fn)
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func Head(path string, fn httprouter.Handle) {
	r.Router.HEAD(path, fn)
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func Options(path string, fn httprouter.Handle) {
	r.Router.OPTIONS(path, fn)
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func Patch(path string, fn httprouter.Handle) {
	r.Router.PATCH(path, fn)
}

// Post is a shortcut for router.Handle("POST", path, handle)
func Post(path string, fn httprouter.Handle) {
	r.Router.POST(path, fn)
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func Put(path string, fn httprouter.Handle) {
	r.Router.PUT(path, fn)
}
