package rarch_encoding

import (
	"bytes"
	"encoding/binary"
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

type Immediate struct {
	length uint8
	value  uint64
}

func (i *Immediate) Encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, i.value)

	bytes := buf.Bytes()
	switch i.length {
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

func DecodeInstruction(instr []byte) Instruction {
	prefix := instr[0]
	opcode := instr[1]
	switch opcode >> 6 {
	case FormatA_Enum:
		return FormatA{prefix: prefix, opcode: opcode}
	case FormatB_Enum:
		return FormatB{prefix: prefix, opcode: opcode, register: instr[2]}
	case FormatC_Enum:
		buf := new(bytes.Buffer)
		buf.Write(instr[3:])
		immval, err := binary.ReadUvarint(buf)
		if err != nil {
			return nil
		}
		imm := Immediate{value: immval, length: ((instr[2] >> 6) & 0x3)}
		return FormatC{prefix: prefix, opcode: opcode, register: instr[2], immediate: imm}
	case FormatD_Enum:
		return FormatD{prefix: prefix, opcode: opcode, destination_register: instr[2], source_register: instr[2]}
	}

	return nil
}

type Instruction interface {
	Encode() []byte
}

type FormatA struct {
	prefix byte
	opcode byte
}

func (i FormatA) Encode() []byte {
	return []byte{i.prefix, i.opcode}
}

type FormatB struct {
	prefix   byte
	opcode   byte
	register byte
}

func (i FormatB) Encode() []byte {
	return []byte{i.prefix, i.opcode, i.register}
}

type FormatC struct {
	prefix    byte
	opcode    byte
	register  byte
	immediate Immediate
}

func (i FormatC) Encode() []byte {
	return append([]byte{i.prefix, i.opcode, i.register}, i.immediate.Encode()...)
}

type FormatD struct {
	prefix               byte
	opcode               byte
	destination_register byte
	source_register      byte
}

func (i FormatD) Encode() []byte {
	return []byte{i.prefix, i.opcode, i.destination_register, i.source_register}
}
