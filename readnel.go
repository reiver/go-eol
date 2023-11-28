package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

func ReadNEL(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(NEL), characterNumber)
	return readthisrune(circumstance, runescanner, nel)
}
