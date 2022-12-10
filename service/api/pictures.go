package api

import (
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetPicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileBytes, err := ioutil.ReadFile("pictures/portrait_20.png")
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(fileBytes)
	return
}
