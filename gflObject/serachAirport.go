package gflObject

import "github.com/garclak/gflgo/gflConst"
import "github.com/garclak/gflgo/gflObject"

type SearchAirport struct {
	gflOject.SearchParam
	ICAO : string
	Country : string
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

func NewSearchAirport(inId int) *SearchAirport {
	lt := new(SearchAirport)
	id := inId
	return lt
}
