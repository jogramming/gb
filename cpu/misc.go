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
