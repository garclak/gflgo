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

	lt.elements["start"] = 10
	lt.elements["logon"] = 20
	lt.elements["pilot"] = 30
	lt.elements["pilotList"] = 35
	lt.elements["pilotAircraft"] = 40
	lt.elements["admin"] = 45
	lt.elements["logList"] = 50
	lt.elements["logReport"] = 60
	lt.elements["waypointList"] = 70
	lt.elements["waypoint"] = 80
	lt.elements["aircraftList"] = 90
	lt.elements["aircraft"] = 100
	lt.elements["airportList"] = 110
	lt.elements["airport"] = 120
	lt.elements["runwayList"] = 130
	lt.elements["runway"] = 140
	lt.elements["flight"] = 150
	lt.elements["flightRoute"] = 160
	lt.elements["routeList"] = 170
	lt.elements["route"] = 180
	lt.elements["exit"] = 190

	return lt
}
