package streamdeck

import (
	"image"
	"log"
	"time"
)

var imageSize = 72
var firstPagePixels = 2583
var secondPagePixels = 2601

func (deck *Deck) SetImage(image *image.RGBA, key int) {
	pixels := make([]byte, imageSize*imageSize*3)

	for i := 0; i < imageSize; i++ {
		rowImage := i * imageSize * 4
		rowPixel := i * imageSize * 3
		for c := 0; c < imageSize; c++ {
			colPosI := (c * 4) + rowImage
			colPosP := (imageSize * 3) + rowPixel - (c * 3) - 1

			pixels[colPosP-2] = image.Pix[colPosI+2]
			pixels[colPosP-1] = image.Pix[colPosI+1]
			pixels[colPosP] = image.Pix[colPosI]
		}
	}

	deck.writeFirstPage(pixels[0:firstPagePixels*3], key)
	deck.writeSecondPage(pixels[firstPagePixels*3:], key)
}

func (deck *Deck) writeFirstPage(pixels []byte, key int) error {
	prefix := []byte{
		0x02, 0x01, 0x01, 0x00, 0x00, (byte)(key + 1), 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x42, 0x4D, 0xF6, 0x3C, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x28, 0x00,
		0x00, 0x00, 0x48, 0x00, 0x00, 0x00, 0x48, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xC0, 0x3C, 0x00, 0x00, 0xC4, 0x0E,
		0x00, 0x00, 0xC4, 0x0E, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	prefix = append(prefix, pixels...)

	data := make([]byte, 8191)
	copy(data, prefix)

	if _, err := deck.device.Write(data, 1*time.Second); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (deck *Deck) writeSecondPage(pixels []byte, key int) error {
	prefix := []byte{0x02, 0x01, 0x02, 0x00, 0x01, (byte)(key + 1)}
	prefix = append(prefix, make([]byte, 10)...)
	prefix = append(prefix, pixels...)

	data := make([]byte, 8191)
	copy(data, prefix)

	if _, err := deck.device.Write(data, 1*time.Second); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
