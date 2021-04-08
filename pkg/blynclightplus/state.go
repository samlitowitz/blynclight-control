package blynclightplus

import "fmt"

type Brightness byte

const (
	Dim  Brightness = 0x1
	Full Brightness = 0x0
)

type FlashRate byte

const (
	NoFlash FlashRate = 0x0
	Low     FlashRate = 0x1 << 4
	Medium  FlashRate = 0x1 << 5
	High    FlashRate = 0x1 << 6
)

type RingRate byte

const (
	Off        RingRate = 0
	Once       RingRate = 0x10
	Continuous RingRate = 0x20
)

type RingTone byte

const (
	None RingTone = iota
	Standard
	_
	_
	_
	_
	_
	_
	_
	_
	Circuit
)

type RingVolume byte

const (
	Mute      RingVolume = 0x80
	MinVolume RingVolume = 0x00
	MaxVolume RingVolume = 0x0A
)

type State struct {
	Red   byte
	Green byte
	Blue  byte

	Brightness Brightness
	FlashRate  FlashRate

	RingRate   RingRate
	RingTone   RingTone
	RingMute   bool
	RingVolume RingVolume
}

func Marshal(s State) ([]byte, error) {
	if s.Brightness != Dim && s.Brightness != Full {
		return nil, fmt.Errorf("brightness must be dim (%d) or full (%d), %d given", Dim, Full, s.Brightness)
	}
	if s.FlashRate != NoFlash && s.FlashRate != Low && s.FlashRate != Medium && s.FlashRate != High {
		return nil, fmt.Errorf("flash rate must be no flash (%d), low (%d), medium (%d), or high (%d), %d given", NoFlash, Low, Medium, High, s.FlashRate)
	}
	if s.RingRate != Off && s.RingRate != Once && s.RingRate != Continuous {
		return nil, fmt.Errorf("ring rate must be off (%d), once (%d), or continuous (%d), %d given", Off, Once, Continuous, s.RingRate)
	}
	if s.RingTone < None || s.RingTone > Circuit {
		return nil, fmt.Errorf("ring tone must be between none (%d) and circuit (%d), %d given", None, Circuit, s.RingTone)
	}
	if s.RingVolume < MinVolume || s.RingVolume > MaxVolume {
		return nil, fmt.Errorf("ring volume must be between min (%d) and max (%d), %d given", MinVolume, MaxVolume, s.RingVolume)
	}
	ringVolume := byte(s.RingVolume)
	if s.RingMute {
		ringVolume = ringVolume | byte(Mute)
	}

	return []byte{
		s.Red,
		s.Blue,
		s.Green,
		byte(s.FlashRate) | byte(s.Brightness), // flash rate and brightness
		byte(s.RingRate) | byte(s.RingTone),    // ring rate and ring tone
		ringVolume,                             // ring volume
		0xff,
		0x22,
	}, nil
}

func Unmarshal(data []byte, state *State) error {
	return nil
}
