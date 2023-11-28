package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

func ReadLF(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(LF), characterNumber)
	return readthisrune(circumstance, runescanner, lf)
}
