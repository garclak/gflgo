package gflConst

type SortC struct {
	elements map[string]int
}

func (l *SortC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewSortC() *SortC {
	lt := new(SortC)
	lt.elements = make(map[string]int)
	lt.elements["asc"] = 10
	lt.elements["desc"] = 20
	return lt
}
