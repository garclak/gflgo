package gflConst

type ObjectC struct {
	elements map[string]int
}

func (l *ObjectC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewObjectC() *ObjectC {
	lt := new(ObjectC)
	lt.elements = make(map[string]int)
	lt.elements["pilot"] = 10
	lt.elements["pilotAircraft"] = 20
	lt.elements["flight"] = 20
	lt.elements["flightRoute"] = 20
	lt.elements["waypoint"] = 20
	lt.elements["airport"] = 20
	lt.elements["aircraft"] = 20
	lt.elements["runway"] = 20
	lt.elements["route"] = 20
	return lt
}
