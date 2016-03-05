// Exercise 1.6: Three color Lissajous GIF to stdout
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.RGBA{0x00, 0x00, 0xFF, 0xFF}}

const (
	blackIndex = 0
	greenIndex = 1
	blueIndex  = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 6     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// cycleProgression = how far through the cycle this value of t is.
			animationProgress := t / (2 * math.Pi)
			cycleProgress := animationProgress - math.Floor(animationProgress)

			// Colour the animation blue when it's 'coming back'.
			if cycleProgress < 0.25 || cycleProgress > 0.75 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
			}
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
