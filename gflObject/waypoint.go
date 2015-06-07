package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type Waypoint struct {
	id : int
	Type : gflConst.WaypointC
	Designation : string
	Altitude : int
	LatHemi : gflConst.HemiC
	LatDeg : int
	LatMin : int
	LatSec : int
	LongHemi :gflConst.HemiC
	LongDeg : int
	LongMin : int
	LongSec : int
	Frequency : float32
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

func NewWaypoint(inId int) *Waypoint {
	lt := new(Waypoint)
	id := inId
	return lt
}

