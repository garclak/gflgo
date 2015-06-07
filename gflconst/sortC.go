package gflConst

type SortC int

const (
    ASC SortC = 400 + iota
    DESC
)

/*
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
	lt.elements["asc"] = 600
	lt.elements["desc"] = 610
	return lt
}
*/