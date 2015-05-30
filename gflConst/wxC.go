package gflConst

type WxC struct {
	elements map[string]int
}

func (l *WxC) Const(ref string) int {
	if ret, ok := l.elements[ref]; ok {
		return ret
	} else {
		return -1
	}
}

func NewWxC() *WxC {
	lt := new(WxC)
	lt.elements = make(map[string]int)
	lt.elements["VFR"] = 10
	lt.elements["MVFR"] = 20
	lt.elements["IFR"] = 20
	lt.elements["CAT-I"] = 20
	lt.elements["CAT-II"] = 20
	lt.elements["CAT-III"] = 20
	return lt
}
