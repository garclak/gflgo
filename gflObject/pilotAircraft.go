package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type PilotAircraft struct {
	id : int
	Rank : gflConst.RankC
	CompletedFlights : int
	DayTotalHours : time.Duration
	NightTotalHours : time.Duration
	WxRating : gflConst.WxC
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

