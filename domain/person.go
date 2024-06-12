package domain

type Person struct {
	PersonID    int
	Age         string
	Name        string
	PhoneNumber string
	City        string
	State       string
	Street1     string
	Street2     string
	ZipCode     string
}

//secondary port
type PersonRepository interface {
	FindByID(ID int) (*Person, error)
}
