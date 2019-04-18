package streamdeck

import "log"

var keys = map[byte]int{
	0x0060: 15,
}

func (deck *Deck) Read() {
	totalKeys := keys[byte(deck.device.Info().Product)]

	for {
		buf, err := deck.device.Read(-1, -1)
		if err != nil {
			log.Println("Error", err)
			return
		}

		kd := buf[1:]

		for i := 0; i < totalKeys; i++ {
			if kd[i] == 0x01 {
				log.Println("PRESSED", i)
			}
		}

		// log.Println("Buf", buf)
	}
}
