package rarch_encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	FormatA_Enum uint8 = iota
	FormatB_Enum
	FormatC_Enum
	FormatD_Enum
)

const (
	ImmediateLen8 uint8 = iota
	ImmediateLen16
	ImmediateLen32
	ImmediateLen64
)
const (
	RegisterLen8 uint8 = iota
	RegisterLen16
	RegisterLen32
	RegisterLen64
)

type Immediate struct {
	Length uint8
	Value  uint64
}

func (i Immediate) String() string {
	return fmt.Sprintf("{%X (length: %d bits)}", i.Value, (1<<i.Length)*8)
}

func (i *Immediate) Encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, i.Value)

	bytes := buf.Bytes()
	switch i.Length {
	case ImmediateLen8:
		return bytes[0:1]
	case ImmediateLen16:
		return bytes[0:2]
	case ImmediateLen32:
		return bytes[0:4]
	case ImmediateLen64:
		return bytes[0:8]

	default:
		return nil
	}
}

func DecodeInstruction(instr []byte) (Instruction, uint64) {

	if len(instr) < 2 {
		return nil, uint64(len(instr))
	}

	prefix := instr[0]
	opcode := instr[1]

	switch opcode >> 6 {
	case FormatA_Enum:
		return FormatA{Prefix: prefix, Opcode: opcode}, 2
	case FormatB_Enum:
		return FormatB{Prefix: prefix, Opcode: opcode, Register: instr[2]}, 3
	case FormatC_Enum:
		buf := new(bytes.Buffer)
		buf.Write(instr[3:])
		var immval uint64
		immsize := DecodeRegister(instr[2]).Register_size
		if (opcode == OP_LD || opcode == OP_STR) && prefix != PREF_LITERAL {
			immsize = ImmediateLen64
		}
		switch immsize {
		case RegisterLen8:
			if len(instr) < 4 {
				return nil, 0
			}
			var immvall uint8
			err := binary.Read(buf, binary.BigEndian, &immvall)
			if err != nil {
				return nil, 0
			}
			immval = uint64(immvall)
		case RegisterLen16:
			if len(instr) < 5 {
				return nil, 0
			}
			var immvall uint16
			err := binary.Read(buf, binary.BigEndian, &immvall)
			if err != nil {
				return nil, 0
			}
			immval = uint64(immvall)
		case RegisterLen32:
			if len(instr) < 7 {
				return nil, 0
			}
			var immvall uint32
			err := binary.Read(buf, binary.BigEndian, &immvall)
			if err != nil {
				return nil, 0
			}
			immval = uint64(immvall)
		case RegisterLen64:
			if len(instr) < 11 {
				return nil, 0
			}
			var immvall uint64
			err := binary.Read(buf, binary.BigEndian, &immvall)
			if err != nil {
				return nil, 0
			}
			immval = uint64(immvall)
		}
		imm := Immediate{Value: uint64(immval), Length: immsize}
		return FormatC{Prefix: prefix, Opcode: opcode, Register: instr[2], Immediate: imm}, 3 + (1 << imm.Length)
	case FormatD_Enum:
		return FormatD{Prefix: prefix, Opcode: opcode, Destination_register: instr[2], Source_register: instr[2]}, 4
	}

	return nil, 0
}

type Instruction interface {
	Encode() []byte
	String() string
}

type FormatA struct {
	Prefix byte
	Opcode byte
}

func (i FormatA) String() string {
	return fmt.Sprintf("{%X %X}", i.Prefix, i.Opcode)
}

func (i FormatA) Encode() []byte {
	return []byte{i.Prefix, i.Opcode}
}

type FormatB struct {
	Prefix   byte
	Opcode   byte
	Register byte
}

func (i FormatB) String() string {
	return fmt.Sprintf("{%X %X %X}", i.Prefix, i.Opcode, i.Register)
}

func (i FormatB) Encode() []byte {
	return []byte{i.Prefix, i.Opcode, i.Register}
}

type FormatC struct {
	Prefix    byte
	Opcode    byte
	Register  byte
	Immediate Immediate
}

func (i FormatC) String() string {
	return fmt.Sprintf("{%X %X %X %s}", i.Prefix, i.Opcode, i.Register, i.Immediate)
}

func (i FormatC) Encode() []byte {
	return append([]byte{i.Prefix, i.Opcode, i.Register}, i.Immediate.Encode()...)
}

type FormatD struct {
	Prefix               byte
	Opcode               byte
	Destination_register byte
	Source_register      byte
}

func (i FormatD) String() string {
	return fmt.Sprintf("{%X %X %X %X}", i.Prefix, i.Opcode, i.Destination_register, i.Source_register)
}

func (i FormatD) Encode() []byte {
	return []byte{i.Prefix, i.Opcode, i.Destination_register, i.Source_register}
}
