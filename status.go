package rarch_encoding

const (
	STATUS_CARRY   uint64 = iota
	STATUS_ZERO           = 1 << iota
	STATUS_CINDEX         = 2 << iota
	STATUS_OFFSIZE        = 6 << iota
	STATUS_PARITY         = 8 << iota
	STATUS_SIGN           = 9 << iota
	STATUS_HALTED         = 10 << iota
)
