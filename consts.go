package eol

const (
	lf = '\u000A'
	cr = '\u000D'
	nl = '\u0085'
	ls = '\u2028'
)

const (
	LF = string(lf)
	CR = string(cr)
	CRLF = CR+LF
	NL = string(nl)
	LS = string(ls)
)
