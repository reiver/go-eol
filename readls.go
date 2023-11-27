package eol

import (
	"io"
)

func ReadLS(runescanner io.RuneScanner) (size int, err error) {
	return readthisrune(runescanner, ls)
}
