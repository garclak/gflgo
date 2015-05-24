package gflConst

type FlightStatusC struct {
	elements map[string]int
}

func (l *FlightStatusC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewFlightStatusC() *FlightStatusC {
	lt := new(FlightStatusC)
	lt.elements = make(map[string]int)
	lt.elements["planning"] = 10
	lt.elements["active"] = 20
	lt.elements["aborted"] = 50
	lt.elements["cancelled"] = 55
	lt.elements["completed"] = 100
	return lt
}
