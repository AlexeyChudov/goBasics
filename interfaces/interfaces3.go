package interfaces

var _ User = &user{}
var _ User = &superUser{}

// var _ User = &user{}  неправильно, т.к. user реализует методы интерфейса через указатели,
// что по сути означает, что он не реализует интерфейс
type user struct { //interanal
	FIO, Adress, Phone string
	isBlocked          bool
}
type superUser struct { //interanal
	FIO, Adress, Phone string
	isBlocked          bool
}

type User interface { //external
	// ChangeAdress(adress string)
	// ChangeFIO(FIO string)
	Block()
}

func (u *user) Block() {
	u.isBlocked = true
}

func (s *superUser) Block() {
	s.isBlocked = true
}

func (u *user) ChangeAdress(adress string) {
	u.Adress = adress
}
func (u *user) ChangeFIO(FIO string) {
	u.FIO = FIO
}

func NewUser(adress, FIO, Phone string) User {
	u := superUser{}
	return &u
}
