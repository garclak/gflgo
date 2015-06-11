package gflObject

import "github.com/garclak/gflgo/gflConst"
import "github.com/garclak/gflgo/gflObject"

type SearchPilot struct {
	gflOject.SearchParam
	Callsign : string
	FirstName : string
    Surname : string
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

func NewSearchPilot(inId int) *SearchPilot {
	lt := new(SearchPilot)
	id := inId
	return lt
}
