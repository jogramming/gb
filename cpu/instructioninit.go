package cpu

func (c *Cpu) AddInstructions() {
	// 0x
	c.addInstruction(0x01, LDBCnn, 3, 12)
	c.addInstruction(0x02, LDBCA, 1, 8)
	c.addInstruction(0x06, LDBn, 2, 8)
	c.addInstruction(0x08, LDnnSP, 3, 20)
	c.addInstruction(0x0a, LDABC, 1, 8)
	c.addInstruction(0x0e, LDCn, 2, 8)
	// 1x
	// 2x
	// 3x
	// 4x
	// 5x
	// 6x
	// 7x
	// 8x
	// 9x
	// ax
	// bx
	// cx
	// dx
	// ex
	// fx
}

func (c *Cpu) addInstruction(opcode uint16, handler func(*Cpu), size int, cycles int) {
	instruction := NewInstruction(handler, size, cycles)
	c.Instructions[opcode] = &instruction
}
