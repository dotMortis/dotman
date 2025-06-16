package bashcmd

import (
	"fmt"
)

type Color string

const (
	Red        Color = "\033[31m"
	Green      Color = "\033[32m"
	colorReset       = "\x1b[0m"
)

type IOWriter struct {
	color Color
}

func (i *IOWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s%s%s", i.color, string(p), colorReset)
	return len(p), nil
}

func NewIOWriter(color Color) *IOWriter {
	return &IOWriter{color: color}
}
