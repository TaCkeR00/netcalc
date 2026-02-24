package ipv4

import "fmt"

const (
	MaxUint8      = 0xFF
	MaxUint32     = 0xFFFFFFFF
	ReservedBitsA = 8
	ReservedBitsB = 16
	ReservedBitsC = 24
)

// Address is exported (Capital 'A')
type Address uint32

// Init replaces your addressInit
func Init(a, b, c, d uint8) Address {
	return Address((uint32(a) << 24) | (uint32(b) << 16) | (uint32(c) << 8) | uint32(d))
}

// Methods attached to the type remain mostly the same, just update the type name
func (ad Address) GetAddressRepresentation() (uint8, uint8, uint8, uint8) {
	a := uint8((ad >> 24) & MaxUint8)
	b := uint8((ad >> 16) & MaxUint8)
	c := uint8((ad >> 8) & MaxUint8)
	d := uint8(ad & MaxUint8)
	return a, b, c, d
}

func (ad *Address) setAddress(a, b, c, d uint8) {
	*ad = Address((uint32(a) << 24) | (uint32(b) << 16) | (uint32(c) << 8) | uint32(d))
}

func (ad Address) String() string {
	a, b, c, d := ad.GetAddressRepresentation()
	return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
}
