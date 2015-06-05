package gflConst

type RankC struct {
	elements map[string]int
}

func (l *RankC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewRankC() *RankC {
	lt := new(RankC)
	lt.elements = make(map[string]int)
	lt.elements["ramp operator"] = 500
	lt.elements["flight engineer"] = 510
	lt.elements["first officer"] = 520
	lt.elements["captain"] = 530
	return lt
}
