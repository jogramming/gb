package cpu

import (
	"fmt"
	"github.com/jonas747/gb/mmu"
	"github.com/jonas747/gb/restime"
	"log"
	"runtime/debug"
	"sync"
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

	Running         bool
	LastInstruction *Instruction
	LastOp          uint16

	BreakMode bool
	BreakChan chan bool
	Waiting   bool

	Counter int64

	LastCycle          restime.Counter
	speedCounter       time.Duration
	speedCounterCycles int
	Speed              float32

	sync.Mutex
}

func (c *Cpu) Reset() {
	c2 := new(Cpu)
	c2.Instructions = c.Instructions
	c2.MMU = c.MMU
	c = c2
}

func (c *Cpu) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic! ", r, string(debug.Stack()))
		}
	}()

	if c.Running {
		log.Println("Cpu allready running !?!?!")
		return
	}
	// fmt.Println("Starting cpu...")
	// fmt.Println("Executing bios")
	c.Running = true
	defer func() {
		c.Running = false
	}()
	c.LastCycle = restime.Now()

	//realLastCycle := restime.Now()
	for {
		if !c.Running {
			break
		}

		//started := restime.Now()
		// now := restime.Now()
		// diff := now.Sub(realLastCycle).Duration()

		// if diff < 1 {
		// 	continue
		// } else {
		// 	//log.Println(diff)
		// }

		//taken := restime.Now().Sub(started).Duration()
		//log.Println("Took ", taken)
		//realLastCycle = now

		c.Cycle()

	}

}

func (c *Cpu) CalcSpeed() {
	now := restime.Now()
	diff := now.Sub(c.LastCycle).Duration()
	c.speedCounter += diff
	c.speedCounterCycles++
	c.LastCycle = now

	if c.speedCounterCycles >= 1000 {
		timePerCycle := c.speedCounter / time.Duration(c.speedCounterCycles)
		c.speedCounterCycles = 0
		c.speedCounter = 0

		if timePerCycle <= 0 {
			timePerCycle = 1
		}
		c.Speed = float32(time.Second / time.Duration(timePerCycle))
		log.Println(c.Speed / 1000)
	}
}

func (c *Cpu) Cycle() {
	c.Lock()

	if c.M > 0 {
		c.M--
	}

	c.CalcSpeed()

	if c.M > 0 {
		c.Unlock()
		return
	}

	c.execOp()
	c.Unlock()

	if c.BreakMode {
		c.Waiting = true
		<-c.BreakChan
		c.Waiting = false
	}
	c.Counter++
}

func (c *Cpu) execOp() {
	sizeMod := 1
	op := uint16(c.MMU.ReadByte(c.PC))
	c.PC++
	if op == 0xcb {
		op2 := c.MMU.ReadByte(c.PC)
		c.PC++
		op = (op << 8) | uint16(op2)
		sizeMod = 2
	}
	c.LastOp = op
	c.LastInstruction = nil
	instruction, ok := c.Instructions[op]
	if !ok {
		log.Printf("Unknown instruction [%X] @ location [%x], stopping...\n", op, c.PC-uint16(sizeMod))
		c.Running = false
		return
	}
	c.LastInstruction = instruction
	handler := instruction.Handler
	if handler == nil {
		log.Println("Unimplemented instruction", instruction)
	}
	handler(c)
	c.PC += uint16(instruction.Size) - uint16(sizeMod)
	c.M += uint8(instruction.Cycles / 4)

	if c.MMU.InBios {
		if c.PC >= 0xff {
			c.MMU.InBios = false
		}
	}
}

type CPUState struct {
	A, B, C, D, E, H, L, F byte // General purpose 8 bit registers(f is flag register)
	PC, SP                 uint16
	LastOp                 uint16
	LastMnemonic           string
	Counter                int64
	Speed                  float32
}

func (c *Cpu) GetState() CPUState {
	// c.Lock()
	// defer c.Unlock()

	mn := "unknown"
	if c.LastInstruction != nil {
		mn = c.LastInstruction.Mnemonic
	}
	return CPUState{
		A:            c.A,
		B:            c.B,
		C:            c.C,
		D:            c.D,
		E:            c.E,
		H:            c.H,
		L:            c.L,
		F:            c.F,
		SP:           c.SP,
		PC:           c.PC,
		LastOp:       c.LastOp,
		LastMnemonic: mn,
		Counter:      c.Counter,
		Speed:        c.Speed,
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

func (c *Cpu) getFlags() (zero, sub, half, carry bool) {
	zf := c.F & FLAGZERO
	sf := c.F & FLAGOPERATION
	hf := c.F & FLAGHALFCARRY
	cf := c.F & FLAGCARRY
	if zf > 0 {
		zero = true
	}
	if sf > 0 {
		sub = true
	}
	if hf > 0 {
		half = true
	}
	if cf > 0 {
		carry = true
	}
	return
}

func (c *Cpu) SetPostBoot() {
	c.MMU.SetPostBoot()

	c.A, c.F = SplitRegisters(1)
	c.F = 0xb0
	c.B, c.C = SplitRegisters(0x0013)
	c.D, c.E = SplitRegisters(0x00d8)
	c.H, c.L = SplitRegisters(0x014d)
	c.SP = 0xfffe

	c.PC = 0x100
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
	Handler  func(*Cpu) // Returns the number of extra cpu cycles used
	Size     int        // size in bytes
	Cycles   int        // Number of cycles used normally
	Op       uint16     // 8 bit if not cb pref
	Mnemonic string
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%X", i.Op)
}
