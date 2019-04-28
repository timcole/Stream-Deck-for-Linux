package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/TimothyCole/Stream-Deck-for-Linux/pkg/streamdeck"
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

	go onReady()

	for _, deck := range decks {
		log.Println(deck)
		if err := deck.Open(); err != nil {
			panic(err)
		}

		go deck.Read()

	UpdateDecks:
		for {
			select {
			case <-waiting:
				break UpdateDecks
			default:
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
		}
	}
}
