package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type Aircraft struct {
	id : int
	Manufacturer : string
	Model : string
	Variant : string
	Type :  gflConst.AircraftC
	VMinOperating : int
	VRotation : int
	VClimb : int
	VFlapsLimit : int
	VLanding : int
	MinRunwayLength : int
	WeightTakeOffMax : int
	WeightEmpty : int
	WeightFuelMax : int
	MaxPAXCount : int
	AltitudeMax : int
	MaxRangeNm : int
	MaxEnduranceMinutes : int
	EngineType : gflConst.EngineC
	EngineCount : int
	FuelCruisePPH : int
	FuelMaxPPH : int
	TotalFlightHours = time.Duration
	MfgDate : time.Time
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

func NewAircraft(inId int) *Aircraft {
	lt := new(Aircraft)
	id := inId
	return lt
}

