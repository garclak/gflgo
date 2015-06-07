package gflObject

import "time"

type FlightLeg struct {
	id : int
	CalcMinutes : time.Duration
	ActMinutes : time.Duration
	ArrDateUTC : time.Time
	PlanAltitude :  int
	PlanAirspeed : int
	PlanFuelWeight : int
	MinFuelWeight : int
	ActFuelWeight : int
	PlanWindMagBearing : int
	PlanWindSpeed : int
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

func NewFlightLeg(inId int) *FlightLeg {
	lt := new(FlightLeg)
	id := inId
	return lt
}

