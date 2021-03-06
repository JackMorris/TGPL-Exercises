// Exercise 1.12: serve random Lissajous animation with specified number of cycles
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Determine the cycles to use, either from the URL or the default value (5).
// Print a lissajous animation to the ResponseWriter.
func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if queryCycles, err := strconv.Atoi(r.URL.Query().Get("cycles")); err == nil {
		if queryCycles > 0 {
			cycles = queryCycles
		}
	}
	lissajous(w, float64(cycles))
}

// Print a lissajous animation with the specified number of cycles to out.
func lissajous(out io.Writer, cycles float64) {

	// Setup animation state.
	var palette = []color.Color{color.White, color.Black}
	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
	)
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 6     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// Draw each frame of the animation.
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
