package main

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"strconv"
)

const (
	BASEDIR = "."
	URL     = "http://vignette4.wikia.nocookie.net/robber-penguin-agency/images/6/6e/Small-mario.png/revision/latest?cb=20150107080404"
)

func GetImage(basedir string, uri string, level float64) (image.Image, error) {
	cacheid := GenCacheId(uri, level)
	pic, err := GetCachedImage(basedir, cacheid)
	if err != nil {
		picraw, err := FetchImage(uri)
		if err != nil {
			return nil, err
		}
		pic = Blur(picraw, level)
		w := GetCacheWriter(basedir, cacheid)
		defer w.Close()
		png.Encode(w, pic)
	}
	return pic, nil
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Query().Get("t")
	level, err := strconv.ParseFloat(r.URL.Query().Get("l"), 64)
	if err != nil {
		fmt.Fprintf(w, "invalid l")
	}

	img, err := GetImage(BASEDIR, uri, level)
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}
	png.Encode(w, img)
}

func main() {
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":8080", nil)
}
