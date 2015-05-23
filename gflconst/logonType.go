package gflconst

type LogonType struct {
	elements map[string]int
}

func (l *LogonType) Type(ref string) int {
	return l.elements[ref]
}

func NewLogonType() *LogonType {
	lt := new(LogonType)
	lt.elements = make(map[string]int)
	lt.elements["normal"] = 10
	lt.elements["admin"] = 20
	return lt
}

