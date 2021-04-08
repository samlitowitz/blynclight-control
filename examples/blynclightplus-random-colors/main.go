package main

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/samlitowitz/embrava-blynclight/pkg/blynclightplus/usb"

	"github.com/google/gousb"
	"github.com/samlitowitz/embrava-blynclight/pkg/blynclightplus"
)

func main() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return desc.Vendor == blynclightplus.Vendor && desc.Product == blynclightplus.Product
	})

	for _, dev := range devs {
		defer dev.Close()
	}

	if err != nil {
		log.Fatalf("OpenDevices(): %v", err)
	}

	if len(devs) == 0 {
		log.Fatalf("found no matching devices for %s:%s", blynclightplus.Vendor, blynclightplus.Product)
	}

	dev := devs[0]
	err = dev.SetAutoDetach(true)
	if err != nil {
		log.Fatalf("SetAutoDetach: %v", err)
	}

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("Config: %v", err)
	}
	defer cfg.Close()

	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Interface: %v", err)
	}
	defer intf.Close()

	block := make([]byte, 3)
	ticker := time.NewTicker(1 * time.Second)
	for {
		<-ticker.C
		_, err := rand.Read(block)
		if err != nil {
			log.Fatal(err)
		}
		_, err = dev.Control(
			gousb.ControlOut|gousb.ControlClass|gousb.ControlInterface,
			byte(usb.SetConfiguration),
			uint16(usb.LightAndSoundConfig),
			0,
			[]byte{block[0], block[1], block[2], 0x02, 0x00, 0x00, 0xff, 0x22},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
