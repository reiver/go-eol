package eol

import (
	"io"
)

func ReadCR(runescanner io.RuneScanner) (size int, err error) {
	const runeNumber = 1
	return readthisrune(runescanner, cr, runeNumber)
}
