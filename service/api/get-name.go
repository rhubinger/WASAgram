package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) GetName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var name string
	name, err := rt.db.GetName()
	if err != nil {
		return
	}
	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte(name))
}
