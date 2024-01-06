package rarch_encoding

const (
	STATUS_CARRY   uint64 = 1 << iota
	STATUS_ZERO           = 1 << 1
	STATUS_INDEX          = 1 << 2
	STATUS_OFFSIZE        = 1 << 6
	STATUS_PARITY         = 1 << 8
	STATUS_SIGN           = 1 << 9
	STATUS_HALTED         = 1 << 10
)
const (
	STATUS_CARRY_BIT   uint64 = iota
	STATUS_ZERO_BIT           = 1
	STATUS_INDEX_BIT          = 2
	STATUS_OFFSIZE_BIT        = 6
	STATUS_PARITY_BIT         = 8
	STATUS_SIGN_BIT           = 9
	STATUS_HALTED_BIT         = 10
)
