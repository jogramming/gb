// Misc insutructions
package cpu

func SwapNibles(b byte) byte {
	out := (b >> 4) | (b << 4)
	return out
}

func (c *Cpu) SwapN(in byte) byte {
	c.F = 0
	out := SwapNibles(in)
	if out == 0 {
		c.F |= FLAGZERO
	}
	return out
}

func SWAPA(c *Cpu) { c.A = c.SwapN(c.A) }
func SWAPB(c *Cpu) { c.B = c.SwapN(c.B) }
func SWAPC(c *Cpu) { c.C = c.SwapN(c.C) }
func SWAPD(c *Cpu) { c.D = c.SwapN(c.D) }
func SWAPE(c *Cpu) { c.E = c.SwapN(c.E) }
func SWAPH(c *Cpu) { c.H = c.SwapN(c.H) }
func SWAPL(c *Cpu) { c.L = c.SwapN(c.L) }
func SWAPHL(c *Cpu) {
	val := c.MMU.ReadByte(CombineRegisters(c.H, c.L))
	out := c.SwapN(val)
	c.MMU.WriteByte(CombineRegisters(c.H, c.L), out)
}

/*
Decimal adjust register A.
This instruction adjusts register A so that the
correct representation of Binary Coded Decimal (BCD)
is obtained

TODO: check for carry
*/
func DAA(c *Cpu) {
	_, sub, _, _ := c.getFlags()
	z := false
	if c.A == 0 {
		z = true
	}

	c.setFlags(z, sub, false, false)

	value := c.A
	value = (((value / 10) % 10) << 4) | (value % 10)
	c.A = value
}

func CPL(c *Cpu) {
	c.A = c.A ^ 0xff
	c.F |= FLAGOPERATION
	c.F |= FLAGHALFCARRY
}

func CCF(c *Cpu) {
	cf := c.F & FLAGCARRY
	cz := c.F & FLAGZERO
	c.F = (cf ^ FLAGCARRY) | cz
}

func SCF(c *Cpu) {
	c.F = (c.F & FLAGZERO) | FLAGCARRY
}

// No operation
func NOP(c *Cpu) {}

// Power down cpu untill interrupt occurs
func HALT(c *Cpu) {}

// Halt CPU and LCD display until button is pressed
func STOP(c *Cpu) {}

// his instruction disables interrupts but not
// immediately. Interrupts are disabled after
// instruction after DI is executed.
func DI(c *Cpu) {}

// Enable interrupts. This intruction enables interrupts
// but not immediately. Interrupts are enabled after
// instruction after EI is executed
func EI() {}

// JR cc,n
// Description:
//   If following condition is true then add n to current
//   address and jump to it:
// Use with:
//   n = one byte signed immediate value
//    cc = NZ, Jump if Z flag is reset.
//    cc = Z,  Jump if Z flag is set.
//    cc = NC, Jump if C flag is reset.
//    cc = C,  Jump if C flag is set.

func (c *Cpu) JRCCN(zeroond bool, flag byte, n byte) {
	doJump := false
	if zeroond {
		if flag == 0 {
			doJump = true
		}
	} else {
		if flag > 0 {
			doJump = true
		}
	}

	if doJump {
		c.PC += uint16(int8(n))
	}
}

func JRNZn(c *Cpu) { c.JRCCN(true, c.F&FLAGZERO, c.MMU.ReadByte(c.PC)) }
func JRZn(c *Cpu)  { c.JRCCN(false, c.F&FLAGZERO, c.MMU.ReadByte(c.PC)) }
func JRNCn(c *Cpu) { c.JRCCN(true, c.F&FLAGCARRY, c.MMU.ReadByte(c.PC)) }
func JRCn(c *Cpu)  { c.JRCCN(false, c.F&FLAGCARRY, c.MMU.ReadByte(c.PC)) }

func CALLnn(c *Cpu) {
	b1 := c.MMU.ReadByte(c.PC)
	b2 := c.MMU.ReadByte(c.PC + 1)
	nextInstruction := c.PC + 2

	c.MMU.WriteWord(c.SP, nextInstruction)
	c.SP -= 2

	combined := uint16(b2<<8) | uint16(b1)
	c.PC = combined - 2
}

func JPnn(c *Cpu) {
	newPC := c.MMU.ReadWord(c.PC)
	c.PC = newPC - 2
}
