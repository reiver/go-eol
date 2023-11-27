package eol

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilRuneScanner = erorr.Error("eol: nil rune-scanner")
)

func errNotEOL(r rune) error {
	return erorr.Errorf("eol: %q (%U) is not an end-of-line character", r, r)
}

func errProblemReadingRune(err error, runeNumber uint64) error {
	return erorr.Errorf("eol: problem reading rune №%d of end-of-line sequence: %w", runeNumber, err)
}

func errProblemUnreadingRune(err error, runeNumber uint64, r rune) error {
	return erorr.Errorf("eol: problem unreading rune №%d (%q (%U)) of end-of-line sequence: %w", runeNumber, r, r, err)
}
