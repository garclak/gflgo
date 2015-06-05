package gflObject

type Pilot struct {
	id : int
	FirstName : string
	Surname : string
	Callsign : string
	PasswordHash : string
	PasswordSalt : string
	Permissions : int //logonC
	MinutesDay : int
	MinutesNight : int
	TotalCompletedFlights : string
}

/*
func (l *LogonC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}
*/

func NewPilot(inId int) *Pilot {
	lt := new(Pilot)
	id := inId
	return lt
}

