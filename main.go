package main

import (
	"github.com/jonas747/gb/cpu"
	"github.com/jonas747/gb/mmu"
)

func main() {
	m := new(mmu.MMU)
	m.Initialize()
	c := &cpu.Cpu{
		MMU:             m,
		Stop:            make(chan bool, 1),
		Instructions:    make(map[uint16]*cpu.Instruction),
		DebuggerEnabled: true,
	}
	c.AddInstructions()
	c.Run()
}
