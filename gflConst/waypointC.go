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
	lt.elements["custom"] = 10
	lt.elements["waypoint"] = 20
	lt.elements["VOR"] = 30
	lt.elements["NDB"] = 40
	lt.elements["airport"] = 50
	return lt
}
