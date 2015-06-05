package gflConst

type EngineC struct {
	elements map[string]int
}

func (l *EngineC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewEngineC() *EngineC {
	lt := new(EngineC)
	lt.elements = make(map[string]int)
	lt.elements["piston"] = 100
	lt.elements["turboprop"] = 110
	lt.elements["turboshaft"] = 120
	lt.elements["jet"] = 130
	lt.elements["turbofan"] = 140
	return lt
}
