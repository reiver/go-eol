package eol

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-eol/ls"
)

// ReadLS tries to read the "\u2028" (i.e., line-separator) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\u2028").
//
// If the character read is not a '\u2028', then ReadLS will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadLS(runescanner)
func ReadLS(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(ls.String), characterNumber)
	return readthisrune(circumstance, runescanner, ls.Rune)
}
