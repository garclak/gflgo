package gflObject

type Airport struct {
	id : int
	ICAO : string
	FIR : string
	Name : string
	Country : string
	Remarks : string
	FreqATIS : float32
	FreqTower : float32
	FreqGround : float32
	FreqApproach : float32
	FreqClearance : float32
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

func NewAirport(inId int) *Airport {
	lt := new(Airport)
	id := inId
	return lt
}

