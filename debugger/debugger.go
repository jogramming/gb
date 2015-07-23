package debugger

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

const (
	Delay = 100
)

type R struct {
	A, B, C, D, E, H, L, F byte // General purpose 8 bit registers(f is flag register)
	PC, SP                 uint16
	LastOp                 uint16
	LastMnemonic           string
}

type Debugger struct {
	Paused          bool
	StepChan        bool
	Registers       R
	NumInstructions int
}

func (d *Debugger) Run() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Error starting debugger: ", err)
	}
	go d.EventWatcher()
}

func (d *Debugger) Draw() {
	//	fmt.Println("Drawing...")
	d.NumInstructions++
	delay := time.NewTimer(time.Millisecond * Delay)
	DrawString("GameBoy Emulator Debugger, LIMIT 100MS Per Instruction. IN BIOS", 1, 1, termbox.ColorDefault, termbox.ColorDefault)
	// Draw the registers

	DrawString("---Registers-----", 0, 3, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("A:", 1, 4, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("B:", 1, 5, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("C:", 1, 6, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("D:", 1, 7, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("E:", 1, 8, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("H:", 1, 9, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("L:", 1, 10, termbox.ColorDefault, termbox.ColorDefault)
	f := d.Registers.F
	DrawString(fmt.Sprintf("FLAGS Z: %d N: %d H: %d C: %d", f>>7, (f>>6)&1, f>>5&1, f>>4&1), 1, 11, termbox.ColorDefault, termbox.ColorDefault)

	DrawString("-------", 1, 12, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("SP:", 1, 13, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("PC:", 1, 14, termbox.ColorDefault, termbox.ColorDefault)

	DrawString(fmt.Sprintf("0x%X", d.Registers.A), 4, 4, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.B), 4, 5, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.C), 4, 6, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.D), 4, 7, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.E), 4, 8, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.H), 4, 9, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.L), 4, 10, termbox.ColorDefault, termbox.ColorDefault)
	//DrawString(fmt.Sprintf("0x%X", d.Registers.F), 4, 11, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.SP), 4, 13, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.Registers.PC), 4, 14, termbox.ColorDefault, termbox.ColorDefault)

	DrawString("-------", 1, 15, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("Last OP: 0x%X %s", d.Registers.LastOp, d.Registers.LastMnemonic), 1, 16, termbox.ColorDefault, termbox.ColorRed)

	DrawString(fmt.Sprintf("Insutructions processed: %d", d.NumInstructions), 20, 5, termbox.ColorDefault, termbox.ColorGreen)

	err := termbox.Flush()
	if err != nil {
		fmt.Println(err)
	}
	<-delay.C
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
}

func (d *Debugger) EventWatcher() {
	for {
		e := termbox.PollEvent()
		if e.Type == termbox.EventKey {
			os.Exit(1)
		}
		fmt.Println("Got event")
	}
}

func DrawString(str string, x, y int, fg, bg termbox.Attribute) {
	for index, r := range str {
		rx := x + index
		termbox.SetCell(rx, y, r, fg, bg)
	}
}
