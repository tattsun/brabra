package main

import (
	"code.google.com/p/graphics-go/graphics"
	"image"
	"image/draw"
)

func Blur(input image.Image, level float64) draw.Image {
	output := image.NewNRGBA(input.Bounds())
	graphics.Blur(output, input, &graphics.BlurOptions{StdDev: level})
	return output
}
