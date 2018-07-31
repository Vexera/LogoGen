package main

import (
	"flag"
	"image/png"
	"net/http"

	"github.com/Vexera/LogoGen"
)

func main() {

	listenPtr := flag.String("listen", ":1337", "The address to listen on")
	flag.Parse()

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

		image, err := LogoGen.CreateLogo(topText[0], bottomText[0])
		if err != nil {
			w.WriteHeader(500)
			return
		}

		png.Encode(w, image)
	})

	panic(http.ListenAndServe(*listenPtr, nil))
}
