//go:generate go run -tags generate gen.go

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

var (
	args = []string{
		"--class=Stream Deck for Linux",
	}
	loadingTitle = "Stream Deck for Linux"
)

func onReady() {
	ui, err := lorca.New(fmt.Sprintf("data:text/html,<html><title>%s</title></html>", loadingTitle), "", 800, 640, args...)
	if err != nil {
		log.Fatalln(err)
	}
	defer ui.Close()

	ui.Bind("ready", func() {
		log.Println("UI is ready...")
	})

	var url string
	if os.Getenv("MODE") != "dev" {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalln(err)
		}
		defer ln.Close()
		go http.Serve(ln, http.FileServer(FS))
		url = ln.Addr().String()
	} else {
		url = "localhost:1234"
	}
	ui.Load(fmt.Sprintf("http://%s", url))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	waiting <- true
}
