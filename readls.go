package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

func ReadLS(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(LS), characterNumber)
	return readthisrune(circumstance, runescanner, ls)
}
