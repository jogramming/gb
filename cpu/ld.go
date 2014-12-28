// All the load/store instructions
package cpu

// 8 bit loads

// LD r, n
// Loads n into reigster r
func LDBn(c *Cpu) { c.B = c.MMU.ReadByte(c.PC) }
func LDCn(c *Cpu) { c.C = c.MMU.ReadByte(c.PC) }
func LDDn(c *Cpu) { c.D = c.MMU.ReadByte(c.PC) }
func LDEn(c *Cpu) { c.E = c.MMU.ReadByte(c.PC) }
func LDHn(c *Cpu) { c.H = c.MMU.ReadByte(c.PC) }
func LDLn(c *Cpu) { c.L = c.MMU.ReadByte(c.PC) }

// LD r1, r2
// put value r2 into r1

// Bx
func LDrrBB(c *Cpu) { c.B = c.B }
func LDrrBC(c *Cpu) { c.B = c.C }
func LDrrBD(c *Cpu) { c.B = c.D }
func LDrrBE(c *Cpu) { c.B = c.E }
func LDrrBH(c *Cpu) { c.B = c.H }
func LDrrBL(c *Cpu) { c.B = c.L }

// Cx
func LDrrCB(c *Cpu) { c.C = c.B }
func LDrrCC(c *Cpu) { c.C = c.C }
func LDrrCD(c *Cpu) { c.C = c.D }
func LDrrCE(c *Cpu) { c.C = c.E }
func LDrrCH(c *Cpu) { c.C = c.H }
func LDrrCL(c *Cpu) { c.C = c.L }

// Dx
func LDrrDB(c *Cpu) { c.D = c.B }
func LDrrDC(c *Cpu) { c.D = c.C }
func LDrrDD(c *Cpu) { c.D = c.D }
func LDrrDE(c *Cpu) { c.D = c.E }
func LDrrDH(c *Cpu) { c.D = c.H }
func LDrrDL(c *Cpu) { c.D = c.L }

// Ex
func LDrrEB(c *Cpu) { c.E = c.B }
func LDrrEC(c *Cpu) { c.E = c.C }
func LDrrED(c *Cpu) { c.E = c.D }
func LDrrEE(c *Cpu) { c.E = c.E }
func LDrrEH(c *Cpu) { c.E = c.H }
func LDrrEL(c *Cpu) { c.E = c.L }

// Hx
func LDrrHB(c *Cpu) { c.H = c.B }
func LDrrHC(c *Cpu) { c.H = c.C }
func LDrrHD(c *Cpu) { c.H = c.D }
func LDrrHE(c *Cpu) { c.H = c.E }
func LDrrHH(c *Cpu) { c.H = c.H }
func LDrrHL(c *Cpu) { c.H = c.L }

// Lx
func LDrrLB(c *Cpu) { c.L = c.B }
func LDrrLC(c *Cpu) { c.L = c.C }
func LDrrLD(c *Cpu) { c.L = c.D }
func LDrrLE(c *Cpu) { c.L = c.E }
func LDrrLH(c *Cpu) { c.L = c.H }
func LDrrLL(c *Cpu) { c.L = c.L }

// xHL
func LDrrAHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.A = c.MMU.ReadByte(addr)
}

func LDrrBHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.B = c.MMU.ReadByte(addr)
}

func LDrrCHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.C = c.MMU.ReadByte(addr)
}

func LDrrDHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.D = c.MMU.ReadByte(addr)
}

func LDrrEHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.E = c.MMU.ReadByte(addr)
}
func LDrrHHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.H = c.MMU.ReadByte(addr)
}
func LDrrLHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.L = c.MMU.ReadByte(addr)
}

// HLx
func LDrrHLB(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.B)
}

func LDrrHLC(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.C)
}

func LDrrHLD(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.D)
}

func LDrrHLE(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.E)
}

func LDrrHLH(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.H)
}

func LDrrHLL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.L)
}

func LDrrHLn(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	n := c.MMU.ReadByte(c.PC)
	c.MMU.WriteByte(addr, n)
}

// LD A, n
func LDAA(c *Cpu) { c.A = c.A }
func LDAB(c *Cpu) { c.A = c.B }
func LDAC(c *Cpu) { c.A = c.C }
func LDAD(c *Cpu) { c.A = c.D }
func LDAE(c *Cpu) { c.A = c.E }
func LDAH(c *Cpu) { c.A = c.H }
func LDAL(c *Cpu) { c.A = c.L }

func LDABC(c *Cpu) {
	addr := CombineRegisters(c.B, c.C)
	c.A = c.MMU.ReadByte(addr)
}

func LDADE(c *Cpu) {
	addr := CombineRegisters(c.D, c.E)
	c.A = c.MMU.ReadByte(addr)
}

func LDAHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.A = c.MMU.ReadByte(addr)
}

func LDAnn(c *Cpu) {
	addr := c.MMU.ReadWord(c.PC)
	c.A = c.MMU.ReadByte(addr)
}

func LDAn(c *Cpu) {
	c.A = c.MMU.ReadByte(c.PC)
}

// LD n, A
func LDBA(c *Cpu) { c.B = c.A }
func LDCA(c *Cpu) { c.C = c.A }
func LDDA(c *Cpu) { c.D = c.A }
func LDEA(c *Cpu) { c.E = c.A }
func LDHA(c *Cpu) { c.H = c.A }
func LDLA(c *Cpu) { c.L = c.A }

func LDBCA(c *Cpu) {
	addr := CombineRegisters(c.B, c.C)
	c.MMU.WriteByte(addr, c.A)
}

func LDDEA(c *Cpu) {
	addr := CombineRegisters(c.D, c.E)
	c.MMU.WriteByte(addr, c.A)
}

func LDHLA(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.A)
}

func LDnnA(c *Cpu) {
	addr := c.MMU.ReadWord(c.PC)
	c.MMU.WriteByte(addr, c.A)
}

// LD A, (c+0xff00)
func LDAIOC(c *Cpu) {
	addr := uint16(c.C) + 0xff00
	c.A = c.MMU.ReadByte(addr)
}

// LD (c+0xff00), A
func LDIOCA(c *Cpu) {
	addr := uint16(c.C) + 0xff00
	c.MMU.WriteByte(addr, c.A)
}

// LDH (n+0xff00), A
func LDIOnA(c *Cpu) {
	addr := uint16(c.MMU.ReadByte(c.PC)) + 0xff00
	c.MMU.WriteByte(addr, c.A)
}

// LDH A, (n+0xff00)
func LDAIOn(c *Cpu) {
	addr := uint16(c.MMU.ReadByte(c.PC)) + 0xff00
	c.A = c.MMU.ReadByte(addr)
}

// LDD A, (HL)
func LDDAHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.A = c.MMU.ReadByte(addr)
	c.L--
	if c.L == 255 {
		c.H--
	}
}

// LDD (HL), A
func LDDHLA(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.A)
	c.L--
	if c.L == 255 {
		c.H--
	}
}

// LDI A, (HL)
func LDIAHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.A = c.MMU.ReadByte(addr)
	c.L++
	if c.L == 0 {
		c.H++
	}
}

// LDI (HL), A
func LDIHLA(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	c.MMU.WriteByte(addr, c.A)
	c.L++
	if c.L == 0 {
		c.H++
	}
}

// end 8 bit loads

// 16 bit loads
// LD n, nn
func LDBCnn(c *Cpu) {
	val := c.MMU.ReadWord(c.PC)
	rb, rc := SplitRegisters(val)
	c.B = rb
	c.C = rc
}

func LDDEnn(c *Cpu) {
	val := c.MMU.ReadWord(c.PC)
	d, e := SplitRegisters(val)
	c.D = d
	c.E = e
}

func LDHLnn(c *Cpu) {
	val := c.MMU.ReadWord(c.PC)
	h, l := SplitRegisters(val)
	c.H = h
	c.L = l
}

func LDSPnn(c *Cpu) {
	c.SP = c.MMU.ReadWord(c.PC)
}

// LD SP, HL
func LDSPHL(c *Cpu) {
	c.SP = CombineRegisters(c.H, c.L)
}

// LDHL SP, n
func LDHLSPn(c *Cpu) {
	val := c.MMU.ReadByte(c.PC)
	sp := c.SP
	result, _, hf, cf := ADD16(sp, uint16(val))
	c.setFlags(false, false, hf, cf)
	h, l := SplitRegisters(result)
	c.L = l
	c.H = h
}

// LD (nn), SP
func LDnnSP(c *Cpu) {
	addr := c.MMU.ReadWord(c.PC)
	c.MMU.WriteWord(addr, c.SP)
}

// PUSH rr
func PUSHAF(c *Cpu) {
	val := CombineRegisters(c.A, c.F)
	c.MMU.WriteWord(c.SP, val)
	c.SP -= 2
}
func PUSHBC(c *Cpu) {
	val := CombineRegisters(c.B, c.C)
	c.MMU.WriteWord(c.SP, val)
	c.SP -= 2
}
func PUSHDE(c *Cpu) {
	val := CombineRegisters(c.D, c.E)
	c.MMU.WriteWord(c.SP, val)
	c.SP -= 2
}
func PUSHHL(c *Cpu) {
	val := CombineRegisters(c.H, c.L)
	c.MMU.WriteWord(c.SP, val)
	c.SP -= 2
}

// POP rr
func POPAF(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	a, f := SplitRegisters(val)
	c.A = a
	c.F = f
}
func POPBC(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	b, cc := SplitRegisters(val)
	c.B = b
	c.C = cc
}
func POPDE(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	d, e := SplitRegisters(val)
	c.D = d
	c.E = e
}
func POPHL(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	h, l := SplitRegisters(val)
	c.H = h
	c.L = l
}
