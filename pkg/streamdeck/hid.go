package streamdeck

import (
	"fmt"

	"github.com/zserge/hid"
)

var (
	vendorID   uint16 = 0x0fd9
	productIDs        = []byte{0x0060}
)

// FindDevices returns an array of StreamDeck devices
func FindDevices() ([]*Deck, error) {
	devices := []*Deck{}

	hid.UsbWalk(func(device hid.Device) {
		info := device.Info()
		if info.Vendor != vendorID {
			return
		}

		for _, product := range productIDs {
			if byte(info.Product) != product {
				return
			}
		}

		devices = append(devices, &Deck{
			device: device,
		})
	})

	if len(devices) == 0 {
		return nil, fmt.Errorf("no elgato streamdeck device found")
	}

	return devices, nil
}

// Open opens and claims the HID
func (deck *Deck) Open() error {
	return deck.device.Open()
}

// Close closes the HID
func (deck *Deck) Close() {
	deck.device.Close()
}
