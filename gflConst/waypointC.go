package gflConst

type WaypointC struct {
	elements map[string]int
}

func (l *WaypointC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewWaypointC() *WaypointC {
	lt := new(WaypointC)
	lt.elements = make(map[string]int)
	lt.elements["custom"] = 1000
	lt.elements["waypoint"] = 1010
	lt.elements["VOR"] = 1020
	lt.elements["NDB"] = 1030
	lt.elements["airport"] = 1040
	return lt
}
