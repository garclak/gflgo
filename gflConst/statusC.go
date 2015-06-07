package gflConst

type StatusC int

const (
    Planning StatusC = 600 + iota
    Active
    Aborted
    Cancelled
    Completed
)

/*
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
	lt.elements["planning"] = 200
	lt.elements["active"] = 210
	lt.elements["aborted"] = 220
	lt.elements["cancelled"] = 230
	lt.elements["completed"] = 240
	return lt
}
*/