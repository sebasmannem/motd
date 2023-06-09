package imgview

/* Copied from:
https://github.com/rdterner/tview/blob/master/main.go
*/

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	termbox "github.com/nsf/termbox-go"
)

var palette color.Palette = []color.Color{
	attrColor(termbox.ColorBlack),
	attrColor(termbox.ColorRed),
	attrColor(termbox.ColorGreen),
	attrColor(termbox.ColorYellow),
	attrColor(termbox.ColorBlue),
	attrColor(termbox.ColorMagenta),
	attrColor(termbox.ColorCyan),
	attrColor(termbox.ColorWhite),
}

func upgrade(palette color.Palette) color.Palette {
	for i := 16; i <= 255; i++ {
		palette = append(palette, attrColor(i))
	}
	return palette
}

type attrColor termbox.Attribute

var base = [...]uint8{0x00, 0x5f, 0x87, 0xaf, 0xd7, 0xff}

func (c attrColor) RGBA() (r, g, b, a uint32) {
	switch termbox.Attribute(c) {
	case termbox.ColorBlack:
		return 0, 0, 0, math.MaxUint16
	case termbox.ColorRed:
		return math.MaxUint16, 0, 0, math.MaxUint16
	case termbox.ColorGreen:
		return 0, math.MaxUint16, 0, math.MaxUint16
	case termbox.ColorYellow:
		return math.MaxUint16, math.MaxUint16, 0, math.MaxUint16
	case termbox.ColorBlue:
		return 0, 0, math.MaxUint16, math.MaxUint16
	case termbox.ColorMagenta:
		return math.MaxUint16, 0, math.MaxUint16, math.MaxUint16
	case termbox.ColorCyan:
		return 0, math.MaxUint16, math.MaxUint16, math.MaxUint16
	case termbox.ColorWhite:
		return math.MaxUint16, math.MaxUint16, math.MaxUint16, math.MaxUint16
	}

	switch {
	case c >= 16 && c <= 231:
		c -= 16
		rgba := color.RGBA{R: base[(c/36)%6], G: base[(c/6)%6], B: base[c%6], A: 0xff}
		return rgba.RGBA()
	case c >= 232 && c <= 255:
		x := uint8(0x08 + 0xa*(c-232))
		rgba := color.RGBA{R: x, G: x, B: x, A: 0xff}
		return rgba.RGBA()
	}

	panic("not found")
}

func fit(w, h, iw, ih int) (sw, sh int) {
	if w >= iw && h >= ih {
		// image is smaller than screen
		return iw, ih
	} else if w >= iw {
		// image is taller than the screen
		return iw * h / ih, h
	} else if h >= ih {
		// image is skinnier than the screen
		return w, ih * w / iw
	} else {
		sw, sh = iw*h/ih, h
		if sw <= w {
			return sw, sh
		}
		return w, ih * w / iw
	}
}

func draw(img image.Image) {
	w, h := termbox.Size()
	iw, ih := img.Bounds().Max.X, img.Bounds().Max.Y
	sw, sh := fit(w, h, iw*2, ih)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x >= sw || y >= sh {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
				continue
			}
			xi := x * img.Bounds().Max.X / sw
			yi := y * img.Bounds().Max.Y / sh
			a := palette.Convert(img.At(xi, yi))
			if at, ok := a.(attrColor); ok {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault,
					termbox.Attribute(at))
			}
		}
	}
	termbox.Flush()
}

func DrawOnConsole(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	if termbox.SetOutputMode(termbox.Output256) == termbox.Output256 {
		palette = upgrade(palette)
	}

	lw, lh := termbox.Size()
	draw(img)
loop:
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		case termbox.EventResize:
			termbox.Flush()
			if w, h := termbox.Size(); lw != w || lh != h {
				lw, lh = w, h
				draw(img)
			}
		}
	}
}
