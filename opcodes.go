package rarch_encoding

/*
 * Opcodes in more readable form in doc/opcodes.odf
 */

const (
	OP_NOP = 0x00
	OP_HLT = 0x01
	OP_UD  = 0x02

	OP_PUSHA = 0x10
	OP_PUSHF = 0x11
	OP_SAVE  = 0x13

	OP_POPA    = 0x20
	OP_POPF    = 0x21
	OP_RESTORE = 0x23

	OP_PUSH = 0x40
	OP_POP  = 0x41

	OP_JMP = 0x50
	OP_JC  = 0x51
	OP_JNC = 0x52
	OP_JE  = 0x53
	OP_JNE = 0x54
	OP_JP  = 0x55
	OP_JNP = 0x56
	OP_JL  = 0x57
	OP_JG  = 0x58
	OP_JEL = 0x59
	OP_JEG = 0x5A

	OP_CALL = 0x60
	OP_RET  = 0x61

	OP_INC = 0x70
	OP_DEC = 0x71

	OP_LD  = 0x80
	OP_STR = 0x81
	OP_LEA = 0x82

	OP_CPY = 0xC0
	OP_SWP = 0xC1

	OP_ADD = 0xD0
	OP_SUB = 0xD1
	OP_MUL = 0xD2
	OP_DIV = 0xD3
	OP_MOD = 0xD4
)

/*
 * Address Prefixes in more readable form in doc/opcodes.odf
 */

const (
	PREF_LITERAL  = 0x00
	PREF_ABSOLUTE = 0x10
	PREF_PCREL    = 0x11
	PREF_INDREL   = 0x12
)
