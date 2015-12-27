package debugger

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jonas747/gb/cpu"
	"github.com/jonas747/gb/restime"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"time"
)

type Debugger struct {
	cpu              *cpu.Cpu
	Paused           bool
	StepChan         bool
	InStepmode       bool
	currentlyWaiting bool
	State            cpu.CPUState

	recentLog []string
	curDelay  time.Duration
}

func (d *Debugger) Run(c *cpu.Cpu) {
	d.cpu = c

	err := termbox.Init()
	if err != nil {
		fmt.Println("Error starting debugger: ", err)
	}

	d.InStepmode = true
	log.SetOutput(d)
	d.cpu.BreakMode = true
	d.cpu.BreakChan = make(chan bool)
	d.curDelay = 100 * time.Millisecond
	log.Println(restime.Overhead())
	go d.EventWatcher()
	go d.Loop()
}

func (d *Debugger) Loop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			d.Draw()
		}
	}
	ticker.Stop()
}

func (d *Debugger) Draw() {
	d.State = d.cpu.GetState()

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	DrawString("GameBoy Emulator Debugger", 1, 1, termbox.ColorDefault, termbox.ColorDefault)
	mode := "Delayed"
	if d.InStepmode {
		mode = "Stepmode"
	} else if d.curDelay <= 0 {
		mode = "Fullspeed"
	}
	DrawString(fmt.Sprintf("Mode: %s, Delay: %dus, Speed: %fkhz", mode, d.curDelay/1000000, d.State.Speed/1000), 1, 2, termbox.ColorDefault, termbox.ColorCyan)
	// Draw the State

	DrawString("---State-----", 0, 3, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("A:", 1, 4, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("B:", 1, 5, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("C:", 1, 6, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("D:", 1, 7, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("E:", 1, 8, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("H:", 1, 9, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("L:", 1, 10, termbox.ColorDefault, termbox.ColorDefault)
	f := d.State.F
	DrawString(fmt.Sprintf("FLAGS Z: %d N: %d H: %d C: %d", f>>7, (f>>6)&1, f>>5&1, f>>4&1), 1, 11, termbox.ColorDefault, termbox.ColorDefault)

	DrawString("-------", 1, 12, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("SP:", 1, 13, termbox.ColorDefault, termbox.ColorDefault)
	DrawString("PC:", 1, 14, termbox.ColorDefault, termbox.ColorDefault)

	DrawString(fmt.Sprintf("0x%X", d.State.A), 4, 4, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.B), 4, 5, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.C), 4, 6, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.D), 4, 7, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.E), 4, 8, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.H), 4, 9, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.L), 4, 10, termbox.ColorDefault, termbox.ColorDefault)
	//DrawString(fmt.Sprintf("0x%X", d.State.F), 4, 11, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.SP), 4, 13, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("0x%X", d.State.PC), 4, 14, termbox.ColorDefault, termbox.ColorDefault)

	DrawString("-------", 1, 15, termbox.ColorDefault, termbox.ColorDefault)
	DrawString(fmt.Sprintf("Last OP: 0x%X %s", d.State.LastOp, d.State.LastMnemonic), 1, 16, termbox.ColorDefault, termbox.ColorRed)

	DrawString(fmt.Sprintf("Insutructions processed: %d", d.State.Counter), 20, 5, termbox.ColorDefault, termbox.ColorGreen)

	d.DrawLog()

	err := termbox.Flush()
	if err != nil {
		log.Println(err)
	}
}

func (d *Debugger) EventWatcher() {
	for {
		e := termbox.PollEvent()
		if e.Type == termbox.EventKey {
			if e.Key == termbox.KeyEsc {
				os.Exit(1)
			} else if e.Key == termbox.KeyArrowUp {
				d.curDelay += 10 * time.Millisecond
			} else if e.Key == termbox.KeyArrowDown {
				d.curDelay -= 10 * time.Millisecond
			} else if e.Key == termbox.KeySpace {
				d.cpu.BreakMode = !d.cpu.BreakMode
			} else if d.cpu.Waiting {
				d.cpu.BreakChan <- true
			}
		}
	}
}

func (d *Debugger) Write(p []byte) (n int, err error) {
	if len(p) > 1 {
		p = p[:len(p)-1]
	}

	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := len(lines) - 1; i >= 0; i-- {
		d.recentLog = append(d.recentLog, lines[i])
	}

	if len(d.recentLog) > 30 {
		newLog := make([]string, 0)
		for i := len(d.recentLog) - 30; i < len(d.recentLog); i++ {
			newLog = append(newLog, d.recentLog[i])
		}
		d.recentLog = newLog
	}
	return len(p), nil
}

func (d *Debugger) DrawLog() {
	count := 0
	for i := len(d.recentLog) - 1; i >= 0; i-- {
		DrawString(d.recentLog[i], 27, 7+count, termbox.ColorDefault, termbox.ColorMagenta)
		count++
	}
}

func DrawString(str string, x, y int, fg, bg termbox.Attribute) {
	for index, r := range str {
		rx := x + index
		termbox.SetCell(rx, y, r, fg, bg)
	}
}
