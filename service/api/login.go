package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !request.Valid() {
		//ctx.Logger.Error("enroll: error validating JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Send the response
	var response = LoginResult{Identifier: ("@" + request.Name)}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
