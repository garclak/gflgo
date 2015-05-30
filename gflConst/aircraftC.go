package gflConst

type AircraftC struct {
	elements map[string]int
}

func (l *AircraftC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewAircraftC() *AircraftC {
	lt := new(AircraftC)
	lt.elements = make(map[string]int)
	lt.elements["glider"] = 10
	lt.elements["light"] = 20
	lt.elements["business"] = 30
	lt.elements["cargo"] = 40
	lt.elements["airliner"] = 50
	lt.elements["military"] = 60
	lt.elements["helicopter"] = 70
	lt.elements["special"] = 80
	return lt
}
