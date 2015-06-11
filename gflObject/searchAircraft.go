package gflObject

import "github.com/garclak/gflgo/gflObject"

type SearchAircraft struct {
	gflOject.SearchParam
	Manufacturer : string
	Designation : string
	Ceiling : int
    RangeNm: int
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

func NewSearchAircraft(inId int) *SearchAircraft {
	lt := new(SearchAircraft)
	id := inId
	return lt
}
