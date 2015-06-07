package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type Flight struct {
	id : int
	Status : gflConst.StatusC
	FlightNotes : string
	WxRating : gflConst.WxC
	OnDateUTC :  time.Time
	OffDateUTC : time.Time
	DurationDay : time.Duration
	DurationNight : time.Duration
	DepCloudCeiling : int
	DepVisDistance : int
	DepOAT : int
	DepWindMagBearing : int
	DepWindSpeed : int
	DepBarometer : int
	ArrCloudCeiling : int
	ArrVisDistance : int
	ArrOAT : int
	ArrWindMagBearing : int
	ArrWindSpeed : int
	ArrBarometer : int
	AltCloudCeiling : int
	AltVisDistance : int
	AltOAT : int
	AltWindMagBearing : int
	AltWindSpeed : int
	AltBarometer : int
	PassengerWeight : int
	CargoWeight : int
	PassengerCount : int
	ZeroFuelWeight : int
	ManifestNotes : string
	MinutesTaxi : int
	MinutesBurn : int
	MinutesCont : int
	MinutesRes : int
	MinutesAlt : int
	MinutesWx : int
	FuelWeightTaxi : int
	FuelWeightBurn : int
	FuelWeightCont : int
	FuelWeightRes : int
	FuelWeightAlt : int
	FuelWeightWx : int
	DebriefReport : string
	IncidentRecorded : bool
	IncidentReport : string
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

func NewFlight(inId int) *Flight {
	lt := new(Aircraft)
	id := inId
	return lt
}

