package writer

type BashCmdWriter interface {
	WriteStd(bytes *[]byte) error
	WriteErr(bytes *[]byte) error
	Close() error
}
