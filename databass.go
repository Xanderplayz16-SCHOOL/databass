package main

import (
//	"binary"
	"encoding/binary"
	"reflect"
)

const MAGIC_NUMBER = 0xE102

func assert(condition bool) {
	if (!condition) {
		panic("Assertion failed.")
	}
}

func _bencode(data interface{}) Buffer { // No, not bittorent.
	buf := Buffer{}
	binary.Encode(buf, binary.BigEndian, data)
	return buf
}

func bencode(data interface{}) Buffer {
	buf := Buffer{}
	encoded := _bencode(data)
	buf.Append(_bencode(len(encoded))...)
	buf.Append(_bencode(101)...) // Type
	buf.Append(encoded...)
	return buf
}

type Buffer []byte

func (this *Buffer) Append(val ...byte) {
	*this = append(*this, val...)
}

func init() {
    binary.Encode()
}
