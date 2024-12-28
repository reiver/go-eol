package eol

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilRuneScanner = erorr.Error("eol: nil rune-scanner")
)

func errNotEOL(r rune) error {
	return erorr.Errorf("eol: %q (%U) is not an end-of-line character", r, r)
}
