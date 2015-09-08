package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	//"log"
	"os"

	"github.com/facchinm/go-serial"
)

var (
	baud   = flag.Int("baud", 4800, "baud rate")
	bits   = flag.Int("bits", 8, "data bits")
	parity = flag.String("parity", "none", "parity bit(none/odd/even/mark/space)")
	stop   = flag.String("stop", "none", "stop bit()")
	raw    = flag.Bool("raw", false, "raw input mode")
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	pb, ok := map[string]serial.Parity{
		"none":  serial.PARITY_NONE,
		"odd":   serial.PARITY_ODD,
		"even":  serial.PARITY_EVEN,
		"mark":  serial.PARITY_MARK,
		"space": serial.PARITY_SPACE,
	}[*parity]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}
	sb, ok := map[string]serial.StopBits{
		"one":          serial.STOPBITS_ONE,
		"onepointfive": serial.STOPBITS_ONEPOINTFIVE,
		"two":          serial.STOPBITS_TWO,
	}[*stop]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	mode := &serial.Mode{
		BaudRate: *baud,
		Parity:   pb,
		DataBits: *bits,
		StopBits: sb,
	}
	port, err := serial.OpenPort(flag.Arg(0), mode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer port.Close()

	go io.Copy(os.Stdout, port)

	if *raw {
		io.Copy(port, os.Stdin)
	} else {
		out := bufio.NewReader(os.Stdin)
		for {
			b, _, err := out.ReadLine()
			if err != nil {
				break
			}
			port.Write(b)
		}
	}
}
