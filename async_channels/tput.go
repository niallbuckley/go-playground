package main

import (
	"fmt"
	"os"
)

type Tput struct{}

func (t Tput) Sc(w *os.File) {
	fmt.Fprint(w, "\0337") // Save cursor position
}

func (t Tput) Rc(w *os.File) {
	fmt.Fprint(w, "\0338") // Restore cursor position
}

func (t Tput) ClearLine(w *os.File) {
	fmt.Fprint(w, "\033[2K") // Clear entire line
}

func (t Tput) Cuu(w *os.File, n int) {
	fmt.Fprintf(w, "\033[%dA", n) // Move cursor up 'n' lines
}

