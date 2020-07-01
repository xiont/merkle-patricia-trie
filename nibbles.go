package main

import (
	"fmt"
)

type Nibble byte

func IsNibble(nibble byte) bool {
	n := int(nibble)
	// 0-9 && a-f
	return n >= 0 && n < 16
}

func FromNibbleByte(n byte) (Nibble, error) {
	if !IsNibble(n) {
		return 0, fmt.Errorf("non-nibble byte: %v", n)
	}
	return Nibble(n), nil
}

// nibbles contain one nibble per byte
func FromNibbleBytes(nibbles []byte) ([]Nibble, error) {
	ns := make([]Nibble, 0, len(nibbles))
	for _, n := range nibbles {
		nibble, err := FromNibbleByte(n)
		if err != nil {
			return nil, fmt.Errorf("contains non-nibble byte: %w", err)
		}
		ns = append(ns, nibble)
	}
	return ns, nil
}

func FromByte(b byte) []Nibble {
	return []Nibble{
		Nibble(byte(b >> 4)),
		Nibble(byte(b % 16)),
	}
}

func FromBytes(bs []byte) []Nibble {
	ns := make([]Nibble, 0, len(bs)*2)
	for _, b := range bs {
		ns = append(ns, FromByte(b)...)
	}
	return ns
}

func FromString(s string) []Nibble {
	return FromBytes([]byte(s))
}

func ToPrefixed(ns []Nibble, isLeafNode bool) []Nibble {
	// create prefix
	var prefixBytes []Nibble
	// odd number of nibbles
	if len(ns)%2 > 0 {
		prefixBytes = []Nibble{1}
	} else {
		// even number of nibbles
		prefixBytes = []Nibble{0, 0}
	}

	// append prefix to all nibble bytes
	prefixed := make([]Nibble, 0, len(prefixBytes)+len(ns))
	prefixed = append(prefixed, prefixBytes...)
	for _, n := range ns {
		prefixed = append(prefixed, Nibble(n))
	}

	// update prefix if is leaf node
	if isLeafNode {
		prefixed[0] += 2
	}

	return prefixed
}

func ToBytes(ns []Nibble) []byte {
	buf := make([]byte, 0, len(ns)/2)

	for i := 0; i < len(ns); i += 2 {
		b := byte(ns[i]<<4) + byte(ns[i+1])
		buf = append(buf, b)
	}

	return buf
}