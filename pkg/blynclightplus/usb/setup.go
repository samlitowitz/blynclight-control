package usb

type BRequest byte

const (
	SetConfiguration BRequest = 9
)

type WValue uint16

const (
	LightAndSoundConfig WValue = 0x0200
)
