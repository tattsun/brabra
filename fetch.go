package main

import (
	"errors"
	"image"
	"image/png"
	"net/http"
)

func FetchImage(uri string) (image.Image, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if res.ContentLength > 5000000 {
		return nil, errors.New("too big image size")
	}

	picture, err := png.Decode(res.Body)
	if err != nil {
		return nil, err
	}
	return picture, nil
}
