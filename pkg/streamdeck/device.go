package streamdeck

import "github.com/zserge/hid"

// Deck is a StreamDeck HID
type Deck struct {
	device hid.Device
}
