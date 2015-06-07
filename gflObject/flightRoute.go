package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type FlightRoute struct {
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

func NewFlightRoute(inId int) *FlightRoute {
	lt := new(FlightRoute)
	id := inId
	return lt
}

