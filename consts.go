package eol

import (
	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
)

const (
	LF   = string(lf.Rune)
	CR   = string(cr.Rune)
	CRLF = CR+LF
	NEL  = string(nel.Rune)
	LS   = string(ls.Rune)
)
