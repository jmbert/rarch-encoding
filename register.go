package rarch_encoding

const (
	RegisterGP uint8 = iota
	RegisterIndex
	RegisterMMove
	RegisterControl
)

type Register struct {
	register_type byte
	register_size byte
	register      byte
}

func (r Register) Encode() byte {
	var encoded byte
	encoded |= r.register & (0xF)
	encoded |= (r.register_size & (0x3)) << 4
	encoded |= (r.register_type & (0x3)) << 6

	return encoded
}

func DecodeRegister(encoded byte) Register {
	var r Register
	r.register = encoded & 0xF
	r.register_size = (encoded >> 4) & 0x3
	r.register_type = (encoded >> 6) & 0x3

	return r
}
