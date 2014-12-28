package cpu

import (
	"github.com/jonas747/gb/mmu"
)

const (
	FLAGZERO      = 0x80
	FLAGOPERATION = 0x40
	FLAGHALFCARRY = 0x20
	FLAGCARRY     = 0X10
)

type Cpu struct {
	ClockM int
	ClockT int

	// Registers
	A, B, C, D, E, H, L, F byte   // General purpose 8 bit registers(f is flag register)
	PC, SP                 uint16 // 16 bit registers
	M, T                   int8   // Clock

	Instructions map[uint16]*Instruction
	MMU          *mmu.MMU
}

func NewInstruction(handler func(*Cpu), size int, cycles int) Instruction {
	return Instruction{
		Handler: handler,
		Size:    size,
		Cycles:  cycles,
	}
}

type Instruction struct {
	Handler func(*Cpu) // Returns the number of extra cpu cycles used
	Size    int        // size in bytes
	Cycles  int        // Number of cycles used normally
}

func (c *Cpu) Reset() {

}

func (c *Cpu) Cycle() {

}

func (c *Cpu) setFlags(zero, sub, half, carry bool) {
	// Reset
	c.F = 0
	if zero {
		c.F |= FLAGZERO
	}

	if sub {
		c.F |= FLAGOPERATION
	}

	if half {
		c.F |= FLAGHALFCARRY
	}

	if carry {
		c.F |= FLAGCARRY
	}
}

func CombineRegisters(a, b byte) uint16 {
	return (uint16(a) << 8) + uint16(b)
}

func SplitRegisters(r uint16) (byte, byte) {
	b := byte(r & 255)
	a := byte(r >> 8)
	return a, b
}
