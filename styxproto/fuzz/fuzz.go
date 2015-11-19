package fuzz

import (
	"bytes"

	"aqwari.net/net/styx/styxproto"
)

// Automated fuzz testing

func Fuzz(data []byte) int {
	d := styxproto.NewDecoder(bytes.NewReader(data))
	for d.Next() {
		if len(d.Messages()) == 0 {
			panic("d.Next read 0 messages without returning an error")
		}
		return 1
	}
	return 0
}
