package cpu

// some helpers
func SUB16(a, b uint16) (result uint16, zf, hf, cf bool) {
	// TODO
	return
}

func SUB8(a uint8, b int) (result uint8, zf, hf, cf bool) {
	ac := int(a)
	// Check half burrow
	if ac&0x0F < b&0x0F {
		hf = true
	}
	// Check burrow
	if ac < b {
		cf = true
	}
	result = uint8(int(a) - b)
	if result == 0 {
		zf = true
	}
	return
}

func SUB88(a uint8, b uint8) (result uint8, zf, hf, cf bool) {
	return SUB8(a, int(b))
}

// 16bit add
func ADD16(a, b uint16) (result uint16, zf, hf, cf bool) {
	// Check half carry
	if ((a&0xfff)+(b&0xfff))&0x1000 > 0 {
		hf = true
	}

	// Check carry
	if int(a)+int(b) > 0xffff {
		cf = true
	}

	result = a + b

	if result == 0 {
		zf = true
	}

	return
}

// 8bit add
func ADD8(a uint8, b int) (result uint8, zf, hf, cf bool) {
	// Check half carry
	ac := int(a)
	if ((ac&0xf)+(b&0xf))&0x10 > 0 {
		hf = true
	}
	// Check carry
	if int(a) > 0xff {
		cf = true
	}
	result = uint8(int(a) + b)
	if result == 0 {
		zf = true
	}
	return
}

func ADD88(a, b uint8) (result uint8, zf, hf, cf bool) {
	return ADD8(a, int(b))
}

// 8 bit alu
// ADD n
func (c *Cpu) ADDr(val uint8) {
	result, zf, hf, cf := ADD88(c.A, val)
	c.setFlags(zf, false, hf, cf)
	c.A = result
}

func ADDAA(c *Cpu)  { c.ADDr(c.A) }
func ADDAB(c *Cpu)  { c.ADDr(c.B) }
func ADDAC(c *Cpu)  { c.ADDr(c.C) }
func ADDAD(c *Cpu)  { c.ADDr(c.D) }
func ADDAE(c *Cpu)  { c.ADDr(c.E) }
func ADDAH(c *Cpu)  { c.ADDr(c.H) }
func ADDAL(c *Cpu)  { c.ADDr(c.L) }
func ADDAHL(c *Cpu) { c.ADDr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func ADDAn(c *Cpu)  { c.ADDr(c.MMU.ReadByte(c.PC)) }

// ADC n
// Add + carry flag
func (c *Cpu) ADCr(val uint8) {
	result, zf, hf, cf := ADD8(c.A, int(val)+int((c.F>>4)&0x1))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCAA(c *Cpu)  { c.ADCr(c.A) }
func ADCAB(c *Cpu)  { c.ADCr(c.B) }
func ADCAC(c *Cpu)  { c.ADCr(c.C) }
func ADCAD(c *Cpu)  { c.ADCr(c.D) }
func ADCAE(c *Cpu)  { c.ADCr(c.E) }
func ADCAH(c *Cpu)  { c.ADCr(c.H) }
func ADCAL(c *Cpu)  { c.ADCr(c.L) }
func ADCAHL(c *Cpu) { c.ADCr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func ADCAn(c *Cpu)  { c.ADCr(c.MMU.ReadByte(c.PC)) }

// SUB n
func (c *Cpu) SUBr(val uint8) {
	result, zf, hf, cf := SUB88(c.A, val)
	c.A = result
	c.setFlags(zf, true, hf, cf)
}

func SUBAA(c *Cpu)  { c.SUBr(c.A) }
func SUBAB(c *Cpu)  { c.SUBr(c.B) }
func SUBAC(c *Cpu)  { c.SUBr(c.C) }
func SUBAD(c *Cpu)  { c.SUBr(c.D) }
func SUBAE(c *Cpu)  { c.SUBr(c.E) }
func SUBAH(c *Cpu)  { c.SUBr(c.H) }
func SUBAL(c *Cpu)  { c.SUBr(c.L) }
func SUBAHL(c *Cpu) { c.SUBr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func SUBAn(c *Cpu)  { c.SUBr(c.MMU.ReadByte(c.PC)) }

// SBC n + carry flag
func (c *Cpu) SBCr(val uint8) {
	flag := int((c.F >> 4) & 0x1)
	result, zf, hf, cf := SUB8(c.A, int(val)+flag)
	c.A = result
	c.setFlags(zf, true, hf, cf)
}

func SBCAA(c *Cpu)  { c.SBCr(c.A) }
func SBCAB(c *Cpu)  { c.SBCr(c.B) }
func SBCAC(c *Cpu)  { c.SBCr(c.C) }
func SBCAD(c *Cpu)  { c.SBCr(c.D) }
func SBCAE(c *Cpu)  { c.SBCr(c.E) }
func SBCAH(c *Cpu)  { c.SBCr(c.H) }
func SBCAL(c *Cpu)  { c.SBCr(c.L) }
func SBCAHL(c *Cpu) { c.SBCr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func SBCAn(c *Cpu)  { c.SBCr(c.MMU.ReadByte(c.PC)) }

// AND n
func (c *Cpu) ANDr(val byte) {
	c.A &= val
	z := false
	if c.A == 0 {
		z = true
	}
	c.setFlags(z, false, true, false)
}

func ANDAA(c *Cpu)  { c.ANDr(c.A) }
func ANDAB(c *Cpu)  { c.ANDr(c.B) }
func ANDAC(c *Cpu)  { c.ANDr(c.C) }
func ANDAD(c *Cpu)  { c.ANDr(c.D) }
func ANDAE(c *Cpu)  { c.ANDr(c.E) }
func ANDAH(c *Cpu)  { c.ANDr(c.H) }
func ANDAL(c *Cpu)  { c.ANDr(c.L) }
func ANDAHL(c *Cpu) { c.ANDr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func ANDAn(c *Cpu)  { c.ANDr(c.MMU.ReadByte(c.PC)) }

// OR n
func (c *Cpu) ORr(val byte) {
	c.A |= val
	z := false
	if c.A == 0 {
		z = true
	}
	c.setFlags(z, false, false, false)
}

func ORAA(c *Cpu)  { c.ORr(c.A) }
func ORAB(c *Cpu)  { c.ORr(c.B) }
func ORAC(c *Cpu)  { c.ORr(c.C) }
func ORAD(c *Cpu)  { c.ORr(c.D) }
func ORAE(c *Cpu)  { c.ORr(c.E) }
func ORAH(c *Cpu)  { c.ORr(c.H) }
func ORAL(c *Cpu)  { c.ORr(c.L) }
func ORAHL(c *Cpu) { c.ORr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func ORAn(c *Cpu)  { c.ORr(c.MMU.ReadByte(c.PC)) }

// XOR n
func (c *Cpu) XORr(val byte) {
	c.A ^= val
	z := false
	if c.A == 0 {
		z = true
	}
	c.setFlags(z, false, false, false)
}

func XORAA(c *Cpu)  { c.XORr(c.A) }
func XORAB(c *Cpu)  { c.XORr(c.B) }
func XORAC(c *Cpu)  { c.XORr(c.C) }
func XORAD(c *Cpu)  { c.XORr(c.D) }
func XORAE(c *Cpu)  { c.XORr(c.E) }
func XORAH(c *Cpu)  { c.XORr(c.H) }
func XORAL(c *Cpu)  { c.XORr(c.L) }
func XORAHL(c *Cpu) { c.XORr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func XORAn(c *Cpu)  { c.XORr(c.MMU.ReadByte(c.PC)) }

// CP  n
func (c *Cpu) CPr(val byte) {
	_, zf, hf, cf := SUB88(c.A, val)
	c.setFlags(zf, true, hf, cf)
}

func CPAA(c *Cpu)  { c.CPr(c.A) }
func CPAB(c *Cpu)  { c.CPr(c.B) }
func CPAC(c *Cpu)  { c.CPr(c.C) }
func CPAD(c *Cpu)  { c.CPr(c.D) }
func CPAE(c *Cpu)  { c.CPr(c.E) }
func CPAH(c *Cpu)  { c.CPr(c.H) }
func CPAL(c *Cpu)  { c.CPr(c.L) }
func CPAHL(c *Cpu) { c.CPr(c.MMU.ReadByte(CombineRegisters(c.H, c.L))) }
func CPAn(c *Cpu)  { c.CPr(c.MMU.ReadByte(c.PC)) }

// INC n
func (c *Cpu) INCr(val byte) byte {
	result, zf, hf, _ := ADD88(val, 1)
	curCflag := c.F & FLAGCARRY
	c.setFlags(zf, false, hf, false)
	c.F |= curCflag
	return result
}

func INCA(c *Cpu) { c.A = c.INCr(c.A) }
func INCB(c *Cpu) { c.B = c.INCr(c.B) }
func INCC(c *Cpu) { c.C = c.INCr(c.C) }
func INCD(c *Cpu) { c.D = c.INCr(c.D) }
func INCE(c *Cpu) { c.E = c.INCr(c.E) }
func INCH(c *Cpu) { c.H = c.INCr(c.H) }
func INCL(c *Cpu) { c.L = c.INCr(c.L) }
func INCHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	curVal := c.MMU.ReadByte(addr)
	c.MMU.WriteByte(addr, c.INCr(curVal))
}

// DEC n
func (c *Cpu) DECr(val byte) byte {
	result, zf, hf, _ := SUB88(val, 1)
	curCflag := c.F & FLAGCARRY
	c.setFlags(zf, true, hf, false)
	c.F |= curCflag
	return result
}

func DECA(c *Cpu) { c.A = c.DECr(c.A) }
func DECB(c *Cpu) { c.B = c.DECr(c.B) }
func DECC(c *Cpu) { c.C = c.DECr(c.C) }
func DECD(c *Cpu) { c.D = c.DECr(c.D) }
func DECE(c *Cpu) { c.E = c.DECr(c.E) }
func DECH(c *Cpu) { c.H = c.DECr(c.H) }
func DECL(c *Cpu) { c.L = c.DECr(c.L) }
func DECHL(c *Cpu) {
	addr := CombineRegisters(c.H, c.L)
	curVal := c.MMU.ReadByte(addr)
	c.MMU.WriteByte(addr, c.DECr(curVal))
}

// 16 bit arithmethic
// ADD HL, n
func (c *Cpu) ADDHLrr(val uint16) {
	result, _, hf, cf := ADD16(CombineRegisters(c.H, c.L), val)

	curZflag := c.F & FLAGZERO
	c.setFlags(false, false, hf, cf)
	c.F |= curZflag
	c.H, c.L = SplitRegisters(result)
}

func ADDHLBC(c *Cpu) { c.ADDHLrr(CombineRegisters(c.B, c.C)) }
func ADDHLDE(c *Cpu) { c.ADDHLrr(CombineRegisters(c.D, c.E)) }
func ADDHLHL(c *Cpu) { c.ADDHLrr(CombineRegisters(c.H, c.L)) }
func ADDHLSP(c *Cpu) { c.ADDHLrr(c.SP) }

// ADD SP, n
// Recheck this later
func ADDSPn(c *Cpu) {
	val := int16(c.MMU.ReadByte(c.PC))
	// Check half carry
	cf, hf := false, false
	if ((c.SP&0xfff)+(uint16(val)&0xfff))&0x1000 > 0 {
		hf = true
	}

	// Check carry
	if int(val)+int(c.SP) > 0xffff {
		cf = true
	}
	c.setFlags(false, false, hf, cf)
	c.SP = uint16(int32(val) + int32(c.SP))
}

//	INC rr
func INCrrBC(c *Cpu) {
	curVal := CombineRegisters(c.B, c.C)
	curVal++
	c.B, c.C = SplitRegisters(curVal)
}
func INCrrDE(c *Cpu) {
	curVal := CombineRegisters(c.D, c.E)
	curVal++
	c.D, c.E = SplitRegisters(curVal)
}
func INCrrHL(c *Cpu) {
	curVal := CombineRegisters(c.H, c.L)
	curVal++
	c.H, c.L = SplitRegisters(curVal)
}
func INCrrSP(c *Cpu) {
	c.SP++
}

//	DEC rr
func DECrrBC(c *Cpu) {
	curVal := CombineRegisters(c.B, c.C)
	curVal--
	c.B, c.C = SplitRegisters(curVal)
}

func DECrrDE(c *Cpu) {
	curVal := CombineRegisters(c.D, c.E)
	curVal--
	c.D, c.E = SplitRegisters(curVal)
}

func DECrrHL(c *Cpu) {
	curVal := CombineRegisters(c.H, c.L)
	curVal--
	c.H, c.L = SplitRegisters(curVal)
}

func DECrrSP(c *Cpu) {
	c.SP--
}
