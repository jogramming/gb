package main

import (
	"flag"
	"github.com/jonas747/gb/cpu"
	"github.com/jonas747/gb/debugger"
	"github.com/jonas747/gb/mmu"
	"io/ioutil"
	"log"
	"runtime"
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	path := flag.Arg(0)
	if path == "" {
		log.Println("No game passed")
		return
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error reading file", err)
		return
	}

	m := new(mmu.MMU)
	m.Initialize()
	m.Rom = file
	c := &cpu.Cpu{
		MMU:          m,
		Instructions: make(map[uint16]*cpu.Instruction),
	}
	c.AddInstructions()

	debug := new(debugger.Debugger)
	debug.Run(c)

	c.SetPostBoot()
	c.Run()

	select {}
}
