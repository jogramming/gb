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

func (c *Cpu) AddAr(val uint8) {
	result, zf, hf, cf := ADD88(c.A, val)
	c.setFlags(zf, false, hf, cf)
	c.A = result
}

func (c *Cpu) AdcAr(val uint8) {
	result, zf, hf, cf := ADD8(c.A, int(val)+int((c.F>>4)&0x1))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func (c *Cpu) SubAr(val uint8) {
	result, zf, hf, cf := SUB88(c.A, val)
	c.A = result
	c.setFlags(zf, true, hf, cf)
}

func (c *Cpu) SbcAr(val uint8) {
	flag := int((c.F >> 4) & 0x1)
	result, zf, hf, cf := SUB8(c.A, int(val)+flag)
	c.A = result
	c.setFlags(zf, true, hf, cf)
}

// 8 bit alu
// ADD A, n
func ADDAA(c *Cpu) { c.AddAr(c.A) }
func ADDAB(c *Cpu) { c.AddAr(c.B) }
func ADDAC(c *Cpu) { c.AddAr(c.C) }
func ADDAD(c *Cpu) { c.AddAr(c.D) }
func ADDAE(c *Cpu) { c.AddAr(c.E) }
func ADDAH(c *Cpu) { c.AddAr(c.H) }
func ADDAL(c *Cpu) { c.AddAr(c.L) }

func ADDAHL(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(CombineRegisters(c.H, c.L)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDAn(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.MMU.ReadByte(c.PC))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

// ADC A, n
// Add + carry flag
func ADCAA(c *Cpu) { c.AdcAr(c.A) }
func ADCAB(c *Cpu) { c.AdcAr(c.B) }
func ADCAC(c *Cpu) { c.AdcAr(c.C) }
func ADCAD(c *Cpu) { c.AdcAr(c.D) }
func ADCAE(c *Cpu) { c.AdcAr(c.E) }
func ADCAH(c *Cpu) { c.AdcAr(c.H) }
func ADCAL(c *Cpu) { c.AdcAr(c.L) }

func ADCAHL(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(CombineRegisters(c.H, c.L))+int((c.F>>4)&0x1))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCAn(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.MMU.ReadByte(c.PC))+int((c.F>>4)&0x1))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

// SUB A, n
func SUBAA(c *Cpu) { c.SubAr(c.A) }
func SUBAB(c *Cpu) { c.SubAr(c.B) }
func SUBAC(c *Cpu) { c.SubAr(c.C) }
func SUBAD(c *Cpu) { c.SubAr(c.D) }
func SUBAE(c *Cpu) { c.SubAr(c.E) }
func SUBAH(c *Cpu) { c.SubAr(c.H) }
func SUBAL(c *Cpu) { c.SubAr(c.L) }

func SUBAHL(c *Cpu) {
	val := c.MMU.ReadByte(CombineRegisters(c.H, c.L))
	result, zf, hf, cf := SUB88(c.A, val)
	c.A = result
	c.setFlags(zf, true, hf, cf)
}

func SUBAn(c *Cpu) {
	val := c.MMU.ReadByte(c.PC)
	result, zf, hf, cf := SUB88(c.A, val)
	c.A = result
	c.setFlags(zf, true, hf, cf)
}
