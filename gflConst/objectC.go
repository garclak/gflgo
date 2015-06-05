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
	lt.elements["pilot"] = 400
	lt.elements["pilotAircraft"] = 410
	lt.elements["flight"] = 420
	lt.elements["flightRoute"] = 430
	lt.elements["waypoint"] = 440
	lt.elements["airport"] = 450
	lt.elements["aircraft"] = 460
	lt.elements["runway"] = 470
	lt.elements["route"] = 480
	return lt
}
