package gflConst

type StateC struct {
	elements map[string]int
}

func (l *StateC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewStateC() *StateC {
	lt := new(StateC)
	lt.elements = make(map[string]int)

	lt.elements["start"] = 700
	lt.elements["logon"] = 701
	lt.elements["pilot"] = 702
	lt.elements["pilotList"] = 703
	lt.elements["pilotAircraft"] = 703
	lt.elements["admin"] = 704
	lt.elements["logList"] = 705
	lt.elements["logReport"] = 706
	lt.elements["waypointList"] = 707
	lt.elements["waypoint"] = 708
	lt.elements["aircraftList"] = 709
	lt.elements["aircraft"] = 710
	lt.elements["airportList"] = 711
	lt.elements["airport"] = 712
	lt.elements["runwayList"] = 713
	lt.elements["runway"] = 714
	lt.elements["flight"] = 715
	lt.elements["flightRoute"] = 716
	lt.elements["routeList"] = 717
	lt.elements["route"] = 718
	lt.elements["exit"] = 719

	return lt
}
