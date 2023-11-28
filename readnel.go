package eol

import (
	"io"
)

func ReadNEL(runescanner io.RuneScanner) (size int, err error) {
	const runeNumber = 1
	return readthisrune(runescanner, nel, runeNumber)
}
