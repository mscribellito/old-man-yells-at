package yeller

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	"image/png"

	"github.com/nfnt/resize"
)

var oldman image.Image = func() image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(rawOldMan))
	m, err := png.Decode(reader)
	if err != nil {
		panic(err)
	}
	return m
}()

// YellAt creates an image with Abe Simpson yelling the target.
func YellAt(target image.Image) image.Image {
	bounds := oldman.Bounds()

	yelled := image.NewRGBA(bounds)
	draw.Draw(yelled, bounds, oldman, image.Point{}, draw.Src)

	at := scaleDown(target)
	draw.Draw(yelled, at.Bounds(), at, image.Point{}, draw.Over)

	return yelled
}

func scaleDown(target image.Image) image.Image {
	s := target.Bounds().Size()
	height := float64(s.Y) * (95 / float64(s.X))
	return resize.Resize(95, uint(height), target, resize.Lanczos3)
}
