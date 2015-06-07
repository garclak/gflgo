package gflObject

import "github.com/garclak/gflgo/gflConst"
import "time"

type RouteInfo struct {
	id : int
	Description : string
	DistanceNm : int
	MinAltitude : int
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

func NewRouteInfo(inId int) *RouteInfo {
	lt := new(RouteInfo)
	id := inId
	return lt
}

