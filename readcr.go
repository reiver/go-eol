package eol

import (
	"io"
)

func ReadCR(runescanner io.RuneScanner) (size int, err error) {
	return readthisrune(runescanner, cr)
}
