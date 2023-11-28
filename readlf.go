package eol

import (
	"io"
)

func ReadLF(runescanner io.RuneScanner) (size int, err error) {
	const runeNumber = 1
	return readthisrune(runescanner, lf, runeNumber)
}
