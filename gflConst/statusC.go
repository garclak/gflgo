package gflConst

type StatusC struct {
	elements map[string]int
}

func (l *StatusC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewStatusC() *StatusC {
	lt := new(StatusC)
	lt.elements = make(map[string]int)
	lt.elements["planning"] = 800
	lt.elements["active"] = 810
	lt.elements["aborted"] = 820
	lt.elements["cancelled"] = 830
	lt.elements["completed"] = 840
	return lt
}
