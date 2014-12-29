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
	c.addInstruction(0x11, LDDEnn, 3, 12)
	c.addInstruction(0x12, LDDEA, 1, 8)
	c.addInstruction(0x16, LDDn, 2, 8)
	c.addInstruction(0x1a, LDADE, 1, 8)
	c.addInstruction(0x1e, LDEn, 2, 8)
	// 2x
	c.addInstruction(0x21, LDHLnn, 3, 12)
	c.addInstruction(0x22, LDIHLA, 1, 8)
	c.addInstruction(0x26, LDHn, 2, 8)
	c.addInstruction(0x2a, LDIAHL, 1, 8)
	c.addInstruction(0x2e, LDAn, 2, 8)
	// 3x
	c.addInstruction(0x31, LDSPnn, 3, 12)
	c.addInstruction(0x32, LDDHLA, 1, 8)
	c.addInstruction(0x36, LDHLn, 2, 12)
	c.addInstruction(0x3a, LDDAHL, 1, 8)
	c.addInstruction(0x3e, LDAn, 2, 8)
	// 4x
	c.addInstruction(0x40, LDBB, 1, 4)
	c.addInstruction(0x41, LDBC, 1, 4)
	c.addInstruction(0x42, LDBD, 1, 4)
	c.addInstruction(0x43, LDBE, 1, 4)
	c.addInstruction(0x44, LDBH, 1, 4)
	c.addInstruction(0x45, LDBL, 1, 4)
	c.addInstruction(0x46, LDBHL, 1, 8)
	c.addInstruction(0x47, LDBA, 1, 4)
	c.addInstruction(0x48, LDCB, 1, 4)
	c.addInstruction(0x49, LDCC, 1, 4)
	c.addInstruction(0x4a, LDCD, 1, 4)
	c.addInstruction(0x4b, LDCE, 1, 4)
	c.addInstruction(0x4c, LDCH, 1, 4)
	c.addInstruction(0x4d, LDCL, 1, 4)
	c.addInstruction(0x4e, LDCHL, 1, 8)
	c.addInstruction(0x4f, LDCA, 1, 4)
	// 5x
	c.addInstruction(0x50, LDDB, 1, 4)
	c.addInstruction(0x51, LDDC, 1, 4)
	c.addInstruction(0x52, LDDD, 1, 4)
	c.addInstruction(0x53, LDDE, 1, 4)
	c.addInstruction(0x54, LDDH, 1, 4)
	c.addInstruction(0x55, LDDL, 1, 4)
	c.addInstruction(0x56, LDDHL, 1, 8)
	c.addInstruction(0x57, LDDA, 1, 4)
	c.addInstruction(0x58, LDEB, 1, 4)
	c.addInstruction(0x59, LDEC, 1, 4)
	c.addInstruction(0x5a, LDED, 1, 4)
	c.addInstruction(0x5b, LDEE, 1, 4)
	c.addInstruction(0x5c, LDEH, 1, 4)
	c.addInstruction(0x5d, LDEL, 1, 4)
	c.addInstruction(0x5e, LDEHL, 1, 8)
	c.addInstruction(0x5f, LDEA, 1, 4)
	// 6x
	c.addInstruction(0x60, LDHB, 1, 4)
	c.addInstruction(0x61, LDHC, 1, 4)
	c.addInstruction(0x62, LDHD, 1, 4)
	c.addInstruction(0x63, LDHE, 1, 4)
	c.addInstruction(0x64, LDHH, 1, 4)
	c.addInstruction(0x65, LDHL, 1, 4)
	c.addInstruction(0x66, LDHHL, 1, 8)
	c.addInstruction(0x67, LDHA, 1, 4)
	c.addInstruction(0x68, LDLB, 1, 4)
	c.addInstruction(0x69, LDLC, 1, 4)
	c.addInstruction(0x6a, LDLD, 1, 4)
	c.addInstruction(0x6b, LDLE, 1, 4)
	c.addInstruction(0x6c, LDLH, 1, 4)
	c.addInstruction(0x6d, LDLL, 1, 4)
	c.addInstruction(0x6e, LDLHL, 1, 8)
	c.addInstruction(0x6f, LDLA, 1, 4)
	// 7x
	c.addInstruction(0x70, LDHLB, 1, 8)
	c.addInstruction(0x71, LDHLC, 1, 8)
	c.addInstruction(0x72, LDHLD, 1, 8)
	c.addInstruction(0x73, LDHLE, 1, 8)
	c.addInstruction(0x74, LDHLH, 1, 8)
	c.addInstruction(0x75, LDHLL, 1, 8)
	// Halt instruction here
	c.addInstruction(0x77, LDHLA, 1, 8)
	c.addInstruction(0x78, LDAB, 1, 4)
	c.addInstruction(0x79, LDAC, 1, 4)
	c.addInstruction(0x7a, LDAD, 1, 4)
	c.addInstruction(0x7b, LDAE, 1, 4)
	c.addInstruction(0x7c, LDAH, 1, 4)
	c.addInstruction(0x7d, LDAL, 1, 4)
	c.addInstruction(0x7e, LDAHL, 1, 8)
	c.addInstruction(0x7f, LDAA, 1, 4)
	// 8x
	// 9x
	// ax
	// bx
	// cx
	c.addInstruction(0xc1, POPBC, 1, 12)
	c.addInstruction(0xc5, PUSHBC, 1, 16)
	// dx
	c.addInstruction(0xd1, POPDE, 1, 12)
	c.addInstruction(0xd5, PUSHDE, 1, 16)
	// ex
	c.addInstruction(0xe0, LDIOnA, 2, 12)
	c.addInstruction(0xe1, POPHL, 1, 12)
	c.addInstruction(0xe2, LDIOCA, 2, 8)
	c.addInstruction(0xe5, PUSHHL, 1, 16)
	c.addInstruction(0xea, LDnnA, 3, 16)
	// fx
	c.addInstruction(0xf0, LDAIOn, 2, 12)
	c.addInstruction(0xf1, POPAF, 1, 12)
	c.addInstruction(0xf2, LDAIOC, 2, 8)
	c.addInstruction(0xc5, PUSHAF, 1, 16)
	c.addInstruction(0xfa, LDAnn, 3, 16)

	// Prefix CB
}

func (c *Cpu) addInstruction(opcode uint16, handler func(*Cpu), size int, cycles int) {
	instruction := NewInstruction(handler, size, cycles)
	c.Instructions[opcode] = &instruction
}
