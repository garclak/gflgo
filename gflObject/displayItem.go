package gflObject

type DisplayItem struct {
	id : int
	data : []string
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

func NewDisplayItem(inId int, inSlice []string) *DisplayItem {
	lt := new(DisplayItem)
	data := inSlice[:]
	id := inId
	return lt
}

func (l *DisplayItem) Data() []string {
	return data
}

func (l *DisplayItem) Id() int {
	return id
}
