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
	lt.elements["piston"] = 10
	lt.elements["turboprop"] = 20
	lt.elements["turboshaft"] = 30
	lt.elements["jet"] = 40
	lt.elements["turbofan"] = 50
	return lt
}
