package main

import (
	"image"
	"image/png"
	"net/http"
)

func FetchImage(uri string) (image.Image, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	picture, err := png.Decode(res.Body)
	if err != nil {
		return nil, err
	}
	return picture, nil
}
