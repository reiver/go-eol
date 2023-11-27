package eol

import (
	"io"
)

func ReadNEL(runescanner io.RuneScanner) (size int, err error) {
	return readthisrune(runescanner, nel)
}
