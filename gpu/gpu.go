package gpu

import (
	"github.com/jonas747/gb/cpu"
	"github.com/jonas747/gb/mmu"
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
	"log"
	"runtime"
	"time"
)

type Flag int

const (
	FlagZero Flag = 0
	FlagOne       = 1
)

var Colors = []color.RGBA{
	color.RGBA{255, 255, 255, 255},
	color.RGBA{180, 180, 180, 255},
	color.RGBA{100, 100, 100, 255},
	color.RGBA{0, 0, 0, 0},
}

type LCDControll struct {
	Control                Flag // Bit 7, true for operating, false for shut down
	WindowTileMapSelect    Flag // Bit 6, The window tile map selected, 0: 9800-$9BFF, 1: $9C00-$9FFF
	WindowDisplay          Flag // Bit 5, 0: off, 1 on
	BGWindowTileDataSelect Flag // Bit 4, 0: 8800-$97FF, 1: $8000-$8FFF <- Same area as OBJ
	BGTileMapDisplaySelect Flag // Bit 3, 0: 9800-$9BFF, 1: $9C00-$9FFF
	SpriteSize             Flag // Bit 2, 0: 8*8, 1: 8*16 (width*height)
	SpriteDisplay          Flag // Bit 1, 0: Off, 1: On
	BGWindowDisplay        Flag // Bit 0, 0: off, 1: On
}

type GPU struct {
	MMU       *mmu.MMU
	CPU       *cpu.Cpu
	window    *sdl.Window
	renderer  *sdl.Renderer
	Control   LCDControll
	BGPalette [4]byte
}

func (g *GPU) Run() {
	g.Initialize()
	ticker := time.NewTicker(116249 * time.Nanosecond) // Roughly that per vert line
	eventCheck := 0
	for {
		<-ticker.C
		eventCheck++
		if eventCheck > 10 {
			eventCheck = 0
			g.CheckEvents()
		}
		if !g.CPU.Running {
			continue
		}

		g.MMU.Lock()
		g.Control = g.GetLCDControl()

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

func (g *GPU) CheckEvents() {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			log.Println("Should quit", t.Timestamp)
		}
	}
}

func (g *GPU) DrawLine(line uint8) {
	if line == 145 {
		// Fire v blank interrupt....somehow
	}
	if line < 145 {
		// Draw the friggging line
		if line%8 == 0 {
			// Draw background tiles! : D
			g.DrawBGRow(line)
		}
	} else if line == 153 {
		g.renderer.Present()
		g.renderer.SetDrawColor(0, 0, 0, 255)
		g.renderer.Clear()
	}
}

func (g *GPU) DrawBGRow(line uint8) {
	for i := 0; i < 32; i++ {
		x := i * 8

	}
}

func (g *GPU) GetTileData(index byte) {

}

func (g *GPU) GetTileIndex(x, y int) byte {
	offset := (y * 32) + x
	if g.Control.BGTileMapDisplaySelect == FlagZero {
		// 9800-9BFF
		// 0x1800-1Bff - videoram start subtracted
		return g.MMU.VideoRam[0x1800+offset]
	} else {
		// 9c00-9FFF
		// 1c00-1FFF - videoram start subtracted
		return g.MMU.VideoRam[0x1c00+offset]
	}
}

func (g *GPU) DrawTile(data []byte, x, y int) {
	height := len(data) / 2
	for i := 0; i < height; i++ {
		// Get the color of each pixel in the tile, which is 8 in width
		b1 := data[(i * 2)]
		b2 := data[(i*2)+1]
		for j := 0; j < 8; j++ {
			colIndex := GetTileColor(b1, b2, uint(i))
			col := Colors[colIndex]
			g.renderer.SetDrawColor(col.R, col.G, col.B, col.A)
			g.renderer.DrawPoint(x+j, y+i)
		}
	}
}

// b1 and b2 = The bytes of the row were at
// example:
//   	    	 b1: 0011 0010
//       		 b2: 0100 0010
// In which case:    0211 0030
func GetTileColor(b1, b2 byte, x uint) byte {
	fb := ((b1 >> x) & 1)
	if x == 0 {
		return fb | ((b2 & 1) << 1)
	} else {
		return fb | (((fb >> x) & 1) << 1)
	}
}

// Convenience method for getting the lcd controll registers
func (g *GPU) GetLCDControl() LCDControll {
	register := g.MMU.IO[0x40]
	control := LCDControll{
		Control:                Flag(register >> 7),
		WindowTileMapSelect:    Flag((register >> 6) & 1),
		WindowDisplay:          Flag((register >> 5) & 1),
		BGWindowTileDataSelect: Flag((register >> 4) & 1),
		BGTileMapDisplaySelect: Flag((register >> 3) & 1),
		SpriteSize:             Flag((register >> 2) & 1),
		SpriteDisplay:          Flag((register >> 1) & 1),
		BGWindowDisplay:        Flag(register & 1),
	}
	return control
}

func (g *GPU) Initialize() {
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_EVERYTHING)
	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	window.SetTitle("747 GB Emulator")
	if err != nil {
		panic(err)
	}
	g.renderer = renderer
	g.window = window

	// Later get from register
	g.BGPalette = [4]byte{0, 1, 2, 3}
}
