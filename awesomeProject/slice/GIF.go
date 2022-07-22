// hello this is lissa package
package GIF

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}, color.RGBA{255, 0, 0, 255}}

const (
	whiteIndex = 0
	blackIndex = 1
	redIndex   = 2
	// 对于palette来说的颜色
)

//func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	if len(os.Args) > 1 && os.Args[1] == "web" {
//		handler := func(w http.ResponseWriter, r *http.Request) {
//			Lissajous(w)
//		}
//		http.HandleFunc("/", handler)
//		log.Fatal(http.ListenAndServe("localhost:8080", nil))
//	}
//	Lissajous(os.Stdout)
//}

func Lissajous(out io.Writer, r *http.Request) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	cycles, _ := strconv.ParseFloat(r.FormValue("cycles"), 64)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			s := rand.Intn(3)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(s))
			//img.SetRGBA64(size+int(x*size+0.5), size+int(y*size+0.5), color.RGBA64{1, 1, 1, 1})
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
