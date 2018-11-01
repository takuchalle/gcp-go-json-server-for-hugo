package ogp

import (
	"encoding/json"
	"errors"
	"net/http"

	ogp "github.com/otiai10/opengraph"
)

func OgpHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if len(url) == 0 {
		http.Error(w, errors.New("url query must be set").Error(), http.StatusInternalServerError)
		return
	}
	og, err := ogp.Fetch(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	og = og.ToAbsURL()
	js, err := json.MarshalIndent(og, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
