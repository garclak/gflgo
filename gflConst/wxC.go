package gflConst

type WxC int

const (
    VFR WxC = 900 + iota
    MVFR
    IFR
    CAT1
    CAT2
    CAT3
)

/*
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
	lt.elements["VFR"] = 1100
	lt.elements["MVFR"] = 1110
	lt.elements["IFR"] = 1120
	lt.elements["CAT-I"] = 1130
	lt.elements["CAT-II"] = 1140
	lt.elements["CAT-III"] = 1150
	return lt
}
*/