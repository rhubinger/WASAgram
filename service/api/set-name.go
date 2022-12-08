package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) SetName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rt.db.SetName("Robert")
	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("Name set to Robert"))
}
