package eol

const (
	lf  = '\u000A'
	cr  = '\u000D'
	nel = '\u0085'
	ls  = '\u2028'
)

const (
	LF   = string(lf)
	CR   = string(cr)
	CRLF = CR+LF
	NEL  = string(nel)
	LS   = string(ls)
)
