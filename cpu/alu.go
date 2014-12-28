package cpu

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
func ADDArA(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.A)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArB(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.B)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArC(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.C)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArD(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.D)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArE(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.E)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArH(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.H)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArL(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.L)
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArHL(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(CombineRegisters(c.H, c.L)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADDArn(c *Cpu) {
	result, zf, hf, cf := ADD88(c.A, c.MMU.ReadByte(c.PC))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

// ADC A, n
func ADCArA(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.A+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArB(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.B+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArC(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.C+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArD(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.D+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArE(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.E+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArH(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.H+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArL(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.L+((c.F>>4)&0x1)))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArHL(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(CombineRegisters(c.H, c.L))+int((c.F>>4)&0x1))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}

func ADCArn(c *Cpu) {
	result, zf, hf, cf := ADD8(c.A, int(c.MMU.ReadByte(c.PC))+int((c.F>>4)&0x1))
	c.A = result
	c.setFlags(zf, false, hf, cf)
}
