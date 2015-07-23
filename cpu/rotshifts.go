package cpu

func RLCA(c *Cpu) {
	highBit := c.A & 0x80
	c.A = c.A<<1 | c.A>>7

	c.F |= (highBit >> 3)
}

// func RLA(c *Cpu) {
// 	highBit := c.A & 0x80
// }
