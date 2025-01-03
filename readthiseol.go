package eol

import (
	"io"

	"github.com/reiver/go-eol/cr"
	"github.com/reiver/go-eol/crlf"
	"github.com/reiver/go-eol/lf"
	"github.com/reiver/go-eol/ls"
	"github.com/reiver/go-eol/nel"
)

// ReadThisEOL tries to read the specified end-of-line sequence.
//
// The end-of-line sequences it supports are:
//
//	line-feed       (LF)  (U+000A) ('\n')
//	carriage-return (CR)  (U+000D) ('\r')
//	carriage-return, line-feed    ("\r\n")
//	next-line       (NEL) (U+0085)
//	line-separator  (LS)  (U+2028)
//
// If successful, ReadThisEOL return the number-of-bytes read (to read in the specified end-of-line sequence).
//
// Example usage:
///
//	size, err := eol.ReadThisEOL(runescanner, eol.CRLF)
func ReadThisEOL(runescanner io.RuneScanner, endofline string) (size int, err error) {
	if nil == runescanner {
		return 0, errNilRuneScanner
	}

	switch endofline {
	case lf.String:
		return ReadLF(runescanner)
	case cr.String:
		return ReadCR(runescanner)
	case crlf.String:
		return ReadCRLF(runescanner)
	case nel.String:
		return ReadNEL(runescanner)
	case ls.String:
		return ReadLS(runescanner)
	default:
		return 0, errUnrecognizedEOL(endofline)
	}
}
