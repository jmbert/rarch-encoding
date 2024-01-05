package rarch_encoding

const (
	RegisterGP uint8 = iota
	RegisterIndex
	RegisterMMove
	RegisterControl
)

const (
	REG_R0 uint8 = iota
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_SP
	REG_BP
)
const (
	REG_I0 uint8 = iota
	REG_I1
	REG_I2
	REG_I3
	REG_I4
	REG_I5
	REG_I6
	REG_I7
	REG_I8
	REG_I9
	REG_I10
	REG_I11
	REG_I12
	REG_I13
	REG_I14
	REG_I15
)
const (
	REG_S0 uint8 = iota
	REG_S1
	REG_S2
	REG_S3
	REG_S4
	REG_S5
	REG_S6
	REG_S7
	REG_D0
	REG_D1
	REG_D2
	REG_D3
	REG_D4
	REG_D5
	REG_D6
	REG_D7
)
const (
	REG_C0 uint8 = iota
	REG_C1
	REG_C2
	REG_C3
	REG_C4
	REG_C5
	REG_C6
	REG_C7
	REG_C8
	REG_C9
	REG_C10
	REG_C11
	REG_C12
	REG_C13
	REG_C14
	REG_C15
)

type Register struct {
	Register_type byte
	Register_size byte
	Register      byte
}

func (r Register) Encode() byte {
	var encoded byte
	encoded |= r.Register & (0xF)
	encoded |= (r.Register_size & (0x3)) << 4
	encoded |= (r.Register_type & (0x3)) << 6

	return encoded
}

func DecodeRegister(encoded byte) Register {
	var r Register
	r.Register = encoded & 0xF
	r.Register_size = (encoded >> 4) & 0x3
	r.Register_type = (encoded >> 6) & 0x3

	return r
}
