package main

import (
	"image/png"
	"net/http"

	"github.com/SilverCory/vexeralogomeme"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(400)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		topText, ok := r.URL.Query()["topText"]
		if !ok {
			topText = []string{""}
		}

		bottomText, ok := r.URL.Query()["bottomText"]
		if !ok {
			bottomText = []string{""}
		}

		image, err := vexeralogomeme.CreateLogo(topText[0], bottomText[0])
		if err != nil {
			w.WriteHeader(500)
			return
		}

		png.Encode(w, image)
	})

	http.ListenAndServe(":1337", nil)
}
