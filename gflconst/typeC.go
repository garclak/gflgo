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
	lt.elements["pilot"] = 900
	lt.elements["pilotAircraft"] = 901
	lt.elements["flight"] = 902
	lt.elements["flightRoute"] = 903
	lt.elements["waypoint"] = 904
	lt.elements["aircraft"] = 905
	lt.elements["airport"] = 906
	lt.elements["runway"] = 907
	lt.elements["route"] = 908
	return lt
}
