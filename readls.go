package eol

import (
	"io"
)

func ReadLS(runescanner io.RuneScanner) (size int, err error) {
	const runeNumber = 1
	return readthisrune(runescanner, ls, runeNumber)
}
