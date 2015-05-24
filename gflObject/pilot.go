package gflObject

type Pilot struct {
	id : int
	firstName : string
	surname : string
	callsign : string
	passwordHash : string
	passwordSalt : string
	permissions : int
	minutesDay : int
	minutesNight : int
	totalCompletedFlights : string
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

func NewPilot(inId int) *Pilot {
	lt := new(Pilot)
	id := inId
	return lt
}

func (l *Pilot) Names() (string, string, string) {
	return firstName, surname, callsign
}

func (l *Pilot) SetNames(inName string, inSurname string, inCallsign string) () {
	firstName := inName
	surname := inSurname
	callsign := inCallsign
}

func (l *Pilot) Password() (string, string, string) {
	return passwordHash, passwordSalt, permissions
}

func (l *Pilot) SetPassword(inHash string, inSalt string, inPerm int) () {
	passwordHash := inHash
	passwordSalt := inSalt
	permissions := inPerm
}

