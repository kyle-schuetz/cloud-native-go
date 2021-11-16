package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/cloud-native-go/key_value_store/core"
	"github.com/gorilla/mux"
)

func KeyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	key := vars["key"]

	value, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = core.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
