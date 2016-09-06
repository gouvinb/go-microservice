package middleware

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// PprofHandler routes the pprof pages using httprouter
func PprofHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PprofHandler called")
	p := context.Get(r, "params").(httprouter.Params)

	switch p.ByName("pprof") {
	case "/cmdline":
		pprof.Cmdline(w, r)
	case "/profile":
		pprof.Profile(w, r)
	case "/symbol":
		pprof.Symbol(w, r)
	default:
		pprof.Index(w, r)
	}
}