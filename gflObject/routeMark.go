package gflObject

type RouteMark struct {
	id : int
	SeqNo : int
	MinAltitude : int
	MagBearing : int
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

func NewRouteMark(inId int) *RouteMark {
	lt := new(RouteMark)
	id := inId
	return lt
}

