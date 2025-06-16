package bashcmd

import (
	"fmt"
)

type Color string

const (
	Red        Color = colorRed
	Green      Color = colorGreen
	colorReset       = "\033[0m"
	colorRed         = "\033[31m"
	colorGreen       = "\033[32m"
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
