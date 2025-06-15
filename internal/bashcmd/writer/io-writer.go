package writer

import "fmt"

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type IOWriter struct{}

func (i *IOWriter) WriteStd(bytes *[]byte) error {
	fmt.Printf("%s%s%s", colorGreen, string(*bytes), colorReset)
	return nil

}

func (i *IOWriter) WriteErr(bytes *[]byte) error {
	fmt.Printf("%s%s%s", colorRed, string(*bytes), colorReset)
	return nil
}

func (i *IOWriter) Close() error {
	return nil
}

func NewIOWriter() *IOWriter {
	return &IOWriter{}
}
