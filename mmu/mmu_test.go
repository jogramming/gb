package mmu

import (
	"testing"
)

var mem *MMU

func init() {
	mem = new(MMU)
	mem.Initialize()
}

func assert(a, b interface{}, t *testing.T) {
	if a != b {
		t.Error(a, "!=", b)
	}
}

func TestWorkingRam(t *testing.T) {
	mem.WriteByte(0xc001, 0xf2)

	ret := mem.ReadByte(0xc001)
	assert(ret, byte(0xf2), t)
}
