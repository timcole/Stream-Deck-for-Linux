package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/TimothyCole/Stream-Deck-for-Linux/pkg/streamdeck"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
	"github.com/nfnt/resize"
)

var done = make(chan bool)

var ui *astilectron.Astilectron
var w *astilectron.Window
var err error

// AppName from ldflags
var AppName string

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

	ui, err = astilectron.New(astilectron.Options{
		AppName:           "Stream Deck for Linux",
		BaseDirectoryPath: os.TempDir(),
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer ui.Close()
	ui.HandleSignals()

	ui.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		astilog.Error("App has crashed")
		return
	})

	if err = ui.Start(); err != nil {
		log.Fatal(err)
		return
	}

	// New tray
	var t = ui.NewTray(&astilectron.TrayOptions{
		Image:   astilectron.PtrStr("./streamdeck.png"),
		Tooltip: astilectron.PtrStr("Stream Deck for Linux"),
	})
	t.Create()
	t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) {
		if w != nil {
			w.Show()
		}
		log.Println(e)
		return
	})

	go func() {
		ui.Wait()
		done <- true
	}()

	go onReady()

	for _, deck := range decks {
		log.Println(deck)
		if err := deck.Open(); err != nil {
			panic(err)
		}

		go deck.Read()
		go func(deck *streamdeck.Deck) {
			for {
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
		}(deck)
	}

	<-done
}
