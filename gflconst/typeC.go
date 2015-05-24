package gflConst

type TypeC struct {
	elements map[string]int
}

func (l *TypeC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewTypeC() *TypeC {
	lt := new(TypeC)
	lt.elements = make(map[string]int)
	lt.elements["pilot"] = 10
	lt.elements["pilotAircraft"] = 20
	lt.elements["flight"] = 30
	lt.elements["flightRoute"] = 40
	lt.elements["waypoint"] = 50
	lt.elements["aircraft"] = 60
	lt.elements["airport"] = 70
	lt.elements["runway"] = 80
	lt.elements["route"] = 90
	return lt
}
