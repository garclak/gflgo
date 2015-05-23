package gflconst

type LogonType struct {
  elements map[string]int{
    "normal" : 1,
    "admin"  : 2
  }
}

func (l LogonType) ConstVal(ref string) int {
  return l.elements[ref]
}
