package gpu

import (
	"github.com/jonas747/gb/mmu"
	"time"
)

type GPU struct {
	MMU *mmu.MMU
}

func (g *GPU) Run() {
	ticker := time.NewTicker(116249 * time.Nanosecond) // Roughly that per vert line
	for {
		<-ticker.C
		g.MMU.Lock()

		curLine := g.MMU.IO[0x44]
		curLine++
		if curLine > 153 {
			curLine = 0
		}
		g.MMU.IO[0x44] = curLine
		g.MMU.Unlock()

		g.DrawLine(curLine)
	}
}

func (g *GPU) DrawLine(line uint8) {
	if line == 145 {
		// Fire v blank interrupt....somehow
	}
	if line < 145 {
		// Draw the friggging line
	}
}

func (g *GPU) Present() {

}
