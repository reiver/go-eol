package eol

import (
	"io"
)

func ReadLF(runescanner io.RuneScanner) (size int, err error) {
	return readthisrune(runescanner, lf)
}
