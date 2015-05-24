package gflConst

type LogonC struct {
	elements map[string]int
}

func (l *LogonC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewLogonC() *LogonC {
	lt := new(LogonC)
	lt.elements = make(map[string]int)
	lt.elements["normal"] = 10
	lt.elements["admin"] = 20
	return lt
}
