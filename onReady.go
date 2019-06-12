package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/TimothyCole/Stream-Deck-for-Linux/pkg/bind"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func onReady() {
	var url string
	if os.Getenv("MODE") != "dev" {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalln(err)
		}
		defer ln.Close()
		go http.Serve(ln, http.FileServer(&assetfs.AssetFS{
			Asset:     bind.Asset,
			AssetDir:  bind.AssetDir,
			AssetInfo: bind.AssetInfo,
			Prefix:    "resources",
		}))
		url = ln.Addr().String()
	} else {
		url = "localhost:1234"
	}

	w, err = ui.NewWindow(fmt.Sprintf("http://%s", url), &astilectron.WindowOptions{
		Title:     &AppName,
		Center:    astilectron.PtrBool(true),
		Width:     astilectron.PtrInt(875),
		MinWidth:  astilectron.PtrInt(875),
		Height:    astilectron.PtrInt(700),
		MinHeight: astilectron.PtrInt(700),
	})
	if err != nil {
		log.Println(err)
		done <- true
		return
	}

	w.Create()

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		var s string
		m.Unmarshal(&s)

		log.Println(s)
		if s == "close" {
			w.Hide()
			return nil
		}
		return nil
	})
	w.SendMessage("hello", func(m *astilectron.EventMessage) {
		// Unmarshal
		var s string
		m.Unmarshal(&s)

		// Process message
		astilog.Debugf("received %s", s)
	})

	w.OpenDevTools()
}
