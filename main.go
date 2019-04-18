package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/TimothyCole/Stream-Deck-for-Linux/pkg/streamdeck"
	"github.com/llgcode/draw2d"
	"github.com/nfnt/resize"
)

var waiting = make(chan bool)

func main() {
	decks, err := streamdeck.FindDevices()
	if err != nil {
		panic(err)
	}

	// Close all our devices on exit
	defer func() {
		for _, deck := range decks {
			deck.Close()
		}
	}()

	draw2d.SetFontFolder("./fonts")
	for _, deck := range decks {
		log.Println(deck)
		if err := deck.Open(); err != nil {
			panic(err)
		}

		go deck.Read()

		for {
			// hello, _ := os.Open("hello.png")
			hello, _ := os.Open("Gopher.png")
			defer hello.Close()
			img, err := png.Decode(hello)
			if err != nil {
				log.Println("Error Decoding Image", err)
			}
			m := resize.Resize(72, 72, img, resize.Lanczos3)

			b := m.Bounds()
			image := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
			draw.Draw(image, image.Bounds(), m, b.Min, draw.Src)

			for i := 0; i < 15; i++ {
				deck.SetImage(image, i)
			}

			time.Sleep(time.Second)
		}
		// i := 0
		// for {
		// 	i++
		// 	dest := image.NewRGBA(image.Rect(0, 0, 72, 72))
		// 	gc := draw2dimg.NewGraphicContext(dest)

		// 	// Background
		// 	gc.BeginPath()
		// 	draw2dkit.Rectangle(gc, 0, 0, 72, 72)
		// 	gc.SetFillColor(color.RGBA{0x1E, 0x90, 0xff, 0xff})
		// 	gc.Fill()

		// 	// Text
		// 	gc.SetFontSize(34)
		// 	gc.SetFontData(draw2d.FontData{Name: "Roboto", Family: draw2d.FontFamilySans, Style: draw2d.FontStyleNormal})
		// 	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
		// 	gc.FillStringAt(fmt.Sprintf("%d", i), 5, 50)

		// 	for i := 0; i < 15; i++ {
		// 		deck.SetImage(dest, i)
		// 	}

		// 	time.Sleep(time.Second)
		// }
	}

	<-waiting
}
