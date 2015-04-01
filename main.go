package main

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strconv"
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

func handler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Query().Get("t")
	level, err := strconv.ParseFloat(r.URL.Query().Get("l"), 64)
	if err != nil {
		fmt.Fprintf(w, "invalid l")
	}

	img, err := GetImage(os.Args[1], uri, level)
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}
	png.Encode(w, img)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
