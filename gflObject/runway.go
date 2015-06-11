package gflObject

import "github.com/garclak/gflgo/gflConst"

type Runway struct {
	id : int
	Designation : string
	Surface : gflConst.SurfaceC
	Length : int
	LocFreq : float32
	LocHeading : int
	LocAltFreq : float32
	LocAltHeading : int
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

func NewRunway(inId int) *Runway {
	lt := new(Runway)
	id := inId
	return lt
}

