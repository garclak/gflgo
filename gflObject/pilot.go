package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type Pilot struct {
	id : int
	FirstName : string
	Surname : string
	Callsign : string
    	HighestRank : gflConst.RankC
    	DOB : time.Time
	PasswordHash : string
	PasswordSalt : string
	Permissions : gflConst.LogonC
    	DayTotalHours : time.Duration
	NightTotalHours : time.Duration
	TotalCompletedFlights : int
    	Remarks : string
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

