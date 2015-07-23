package cpu

// duplicated
// func RLCA(c *Cpu) {
// 	highBit := c.A & 0x80
// 	c.A = c.A<<1 | c.A>>7

// 	z := false
// 	if c.A == 0 {
// 		z = true
// 	}

// 	c.setFlags(z, false, false, false)
// 	c.F |= (highBit >> 3)
// }

// func RLA(c *Cpu) {
// 	highBit := c.A & 0x80
// 	val := c.A << 1
// 	lb := c.F & FLAGCARRY
// 	lb = lb >> 4
// 	cf := highBit >> 3

// 	val |= lb
// 	c.A = val

// 	z := false
// 	if c.A == 0 {
// 		z = true
// 	}
// 	c.setFlags(z, false, false, false)
// 	c.F &= cf
// }

func RRCA(c *Cpu) {
	low := c.A & 0x1
	c.A = c.A>>1 | c.A<<7

	z := false
	if c.A == 0 {
		z = true
	}

	c.setFlags(z, false, false, false)
	c.F |= low << 4
}

func RRA(c *Cpu) {
	low := c.A & 0x1
	val := c.A >> 1
	hb := c.F & FLAGCARRY
	hb = hb << 3
	cf := low << 4

	val |= hb
	c.A = val

	z := false
	if c.A == 0 {
		z = true
	}
	c.setFlags(z, false, false, false)
	c.F &= cf
}

func (c *Cpu) RLCn(reg byte) byte {
	hBit := reg & 0x80
	reg = (reg << 1) | (reg >> 7)

	z := false
	if reg == 0 {
		z = true
	}

	c.setFlags(z, false, false, false)
	c.F |= hBit >> 3
	return reg
}

func RLCA(c *Cpu) { c.A = c.RLCn(c.A) }
func RLCB(c *Cpu) { c.B = c.RLCn(c.B) }
func RLCC(c *Cpu) { c.C = c.RLCn(c.C) }
func RLCD(c *Cpu) { c.D = c.RLCn(c.D) }
func RLCE(c *Cpu) { c.E = c.RLCn(c.E) }
func RLCH(c *Cpu) { c.H = c.RLCn(c.H) }
func RLCL(c *Cpu) { c.L = c.RLCn(c.L) }

func RLCHL(c *Cpu) {
	b := c.MMU.ReadByte(CombineRegisters(c.H, c.L))
	b = c.RLCn(b)
	c.MMU.WriteByte(CombineRegisters(c.H, c.L), b)
}

func (c *Cpu) RLn(reg byte) byte {
	highBit := reg & 0x80
	val := reg << 1
	lb := c.F & FLAGCARRY
	lb = lb >> 4
	cf := highBit >> 3

	val |= lb

	z := false
	if reg == 0 {
		z = true
	}
	c.setFlags(z, false, false, false)
	c.F &= cf
	return val
}

func RLA(c *Cpu) { c.A = c.RLn(c.A) }
func RLB(c *Cpu) { c.B = c.RLn(c.B) }
func RLC(c *Cpu) { c.C = c.RLn(c.C) }
func RLD(c *Cpu) { c.D = c.RLn(c.D) }
func RLE(c *Cpu) { c.E = c.RLn(c.E) }
func RLH(c *Cpu) { c.H = c.RLn(c.H) }
func RLL(c *Cpu) { c.L = c.RLn(c.L) }

func RLHL(c *Cpu) {
	b := c.MMU.ReadByte(CombineRegisters(c.H, c.L))
	b = c.RLn(b)
	c.MMU.WriteByte(CombineRegisters(c.H, c.L), b)
}

func BIT7H(c *Cpu) {
	res := c.H >> 7
	cf := (c.F & FLAGCARRY) | FLAGHALFCARRY
	if res <= 0 {
		cf |= FLAGZERO
	}
	c.F = cf
}
