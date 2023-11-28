package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

func ReadCR(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(CR), characterNumber)
	return readthisrune(circumstance, runescanner, cr)
}
