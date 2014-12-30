package cpu

import (
	"fmt"
	"github.com/jonas747/gb/mmu"
	"time"
)

const (
	FLAGZERO      = 0x80
	FLAGOPERATION = 0x40
	FLAGHALFCARRY = 0x20
	FLAGCARRY     = 0X10
)

type Cpu struct {
	// Registers
	A, B, C, D, E, H, L, F byte   // General purpose 8 bit registers(f is flag register)
	PC, SP                 uint16 // 16 bit registers
	M                      uint8  // Machine cycles untill next instruction(clock cycles / 4)

	Instructions map[uint16]*Instruction
	MMU          *mmu.MMU

	Running bool
	Stop    chan bool
}

func (c *Cpu) Reset() {
	c2 := new(Cpu)
	c2.Instructions = c.Instructions
	c2.MMU = c.MMU
	c2.Stop = c.Stop
	c = c2
}

func (c *Cpu) Run() {
	if c.Running {
		fmt.Println("Cpu allready running !?!?!")
		return
	}
	fmt.Println("Starting cpu...")
	fmt.Println("Executing bios")
	c.Running = true
	defer func() {
		c.Running = false
	}()
	ticker := time.NewTicker(time.Duration(952))
	for {
		select {
		case <-ticker.C:
			c.Cycle()
		case <-c.Stop:
			return
		}
	}
}

func (c *Cpu) Cycle() {
	c.M--
	if c.M > 0 {
		return
	}

	c.execOp()
}

func (c *Cpu) execOp() {
	sizeMod := 1
	op := uint16(c.MMU.ReadByte(c.PC))
	c.PC++
	if op == 0xcb {
		op2 := c.MMU.ReadByte(c.PC)
		c.PC++
		op = (op << 8) + uint16(op2)
		sizeMod = 2
	}
	instruction, ok := c.Instructions[uint16(op)]
	if !ok {
		fmt.Printf("Unknown instruction [%X] @ location [%x], stopping...\n", op, c.PC-uint16(sizeMod))
		c.Stop <- true
		return
	}
	handler := instruction.Handler
	if handler == nil {
		fmt.Println("Unimplemented instruction", instruction)
	}
	handler(c)
	fmt.Println("Executed instruction ", instruction)
	c.PC += uint16(instruction.Size) - uint16(sizeMod)
	c.M += uint8(instruction.Cycles / 4)

	if c.MMU.InBios {
		if op >= 0xff {
			c.MMU.InBios = false
		}
	}
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
	Op      uint16
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%X", i.Op)
}
