package eol

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-eol/lf"
)

// ReadLF tries to read the "\n" (i.e., line-feed) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\n").
//
// If the character read is not a '\n', then ReadLF will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadLF(runescanner)
func ReadLF(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(lf.String), characterNumber)
	return readthisrune(circumstance, runescanner, lf.Rune)
}
