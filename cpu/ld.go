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
func LDBB(c *Cpu) { c.B = c.B }
func LDBC(c *Cpu) { c.B = c.C }
func LDBD(c *Cpu) { c.B = c.D }
func LDBE(c *Cpu) { c.B = c.E }
func LDBH(c *Cpu) { c.B = c.H }
func LDBL(c *Cpu) { c.B = c.L }

// Cx
func LDCB(c *Cpu) { c.C = c.B }
func LDCC(c *Cpu) { c.C = c.C }
func LDCD(c *Cpu) { c.C = c.D }
func LDCE(c *Cpu) { c.C = c.E }
func LDCH(c *Cpu) { c.C = c.H }
func LDCL(c *Cpu) { c.C = c.L }

// Dx
func LDDB(c *Cpu) { c.D = c.B }
func LDDC(c *Cpu) { c.D = c.C }
func LDDD(c *Cpu) { c.D = c.D }
func LDDE(c *Cpu) { c.D = c.E }
func LDDH(c *Cpu) { c.D = c.H }
func LDDL(c *Cpu) { c.D = c.L }

// Ex
func LDEB(c *Cpu) { c.E = c.B }
func LDEC(c *Cpu) { c.E = c.C }
func LDED(c *Cpu) { c.E = c.D }
func LDEE(c *Cpu) { c.E = c.E }
func LDEH(c *Cpu) { c.E = c.H }
func LDEL(c *Cpu) { c.E = c.L }

// Hx
func LDHB(c *Cpu) { c.H = c.B }
func LDHC(c *Cpu) { c.H = c.C }
func LDHD(c *Cpu) { c.H = c.D }
func LDHE(c *Cpu) { c.H = c.E }
func LDHH(c *Cpu) { c.H = c.H }
func LDHL(c *Cpu) { c.H = c.L }

// Lx
func LDLB(c *Cpu) { c.L = c.B }
func LDLC(c *Cpu) { c.L = c.C }
func LDLD(c *Cpu) { c.L = c.D }
func LDLE(c *Cpu) { c.L = c.E }
func LDLH(c *Cpu) { c.L = c.H }
func LDLL(c *Cpu) { c.L = c.L }

// xHL
func LDAHL(c *Cpu) { c.A = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }
func LDBHL(c *Cpu) { c.B = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }
func LDCHL(c *Cpu) { c.C = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }
func LDDHL(c *Cpu) { c.D = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }
func LDEHL(c *Cpu) { c.E = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }
func LDHHL(c *Cpu) { c.H = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }
func LDLHL(c *Cpu) { c.L = c.MMU.ReadByte(CombineRegisters(c.H, c.L)) }

// HLx
func LDHLB(c *Cpu) { c.MMU.WriteByte(CombineRegisters(c.H, c.L), c.B) }
func LDHLC(c *Cpu) { c.MMU.WriteByte(CombineRegisters(c.H, c.L), c.C) }
func LDHLD(c *Cpu) { c.MMU.WriteByte(CombineRegisters(c.H, c.L), c.D) }
func LDHLE(c *Cpu) { c.MMU.WriteByte(CombineRegisters(c.H, c.L), c.E) }
func LDHLH(c *Cpu) { c.MMU.WriteByte(CombineRegisters(c.H, c.L), c.H) }
func LDHLL(c *Cpu) { c.MMU.WriteByte(CombineRegisters(c.H, c.L), c.L) }

func LDHLn(c *Cpu) {
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

func LDABC(c *Cpu) { c.A = c.MMU.ReadByte(CombineRegisters(c.B, c.C)) }
func LDADE(c *Cpu) { c.A = c.MMU.ReadByte(CombineRegisters(c.D, c.E)) }
func LDAnn(c *Cpu) { c.A = c.MMU.ReadByte(c.MMU.ReadWord(c.PC)) }
func LDAn(c *Cpu)  { c.A = c.MMU.ReadByte(c.PC) }

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
	c.B, c.C = SplitRegisters(val)
}

func LDDEnn(c *Cpu) {
	val := c.MMU.ReadWord(c.PC)
	c.D, c.E = SplitRegisters(val)
}

func LDHLnn(c *Cpu) {
	val := c.MMU.ReadWord(c.PC)
	c.H, c.L = SplitRegisters(val)
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
	c.H, c.L = SplitRegisters(result)
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
	c.A, c.F = SplitRegisters(val)
}
func POPBC(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	c.B, c.C = SplitRegisters(val)
}
func POPDE(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	c.D, c.E = SplitRegisters(val)
}
func POPHL(c *Cpu) {
	c.SP += 2
	val := c.MMU.ReadWord(c.SP)
	c.H, c.L = SplitRegisters(val)
}
