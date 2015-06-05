package gflConst

type YesNoC struct {
	elements map[string]int
}

func (l *YesNoC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewYesNoC() *YesNoC {
	lt := new(YesNoC)
	lt.elements = make(map[string]int)
	lt.elements["yes"] = 1200
	lt.elements["no"] = 1210
	return lt
}
