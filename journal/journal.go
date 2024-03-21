package journal

import (
	"io"
	"os"
)

func NewWriter(name string) io.Writer {
	// TODO: implement
	_ = name
	return os.Stdout
}
