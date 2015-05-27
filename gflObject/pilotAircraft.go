package gflObject

type PilotAircraft struct {
	id : int
	Rank : string
	CompletedFlights : string
	MinutesDay : int
	MinutesNight : int
	WxCategoryLimit : int //Todo: Const
	TotalCargoWeight : int
	TotalPAXWeight : int
	TotalPAXCount : int
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

func NewPilotAircraft(inId int) *PilotAircraft {
	lt := new(PilotAircraft)
	id := inId
	return lt
}

