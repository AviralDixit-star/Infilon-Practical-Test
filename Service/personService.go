package service

import (
	"log"

	"github.com/AviralDixit-star/Infilon-Practical-Test/domain"
)

//Primary Port
type PersonService interface {
	GetPersonByID(id int) (*domain.Person, error)
}

// Service Implementation
type DefaultPersonService struct {
	//dependency injection
	Repo domain.PersonRepository
}

//receiver function
func (d DefaultPersonService) GetPersonByID(id int) (*domain.Person, error) {
	person, err := d.Repo.FindByID(id)
	if err != nil {
		log.Fatal(err)
		return &domain.Person{}, err
	}
	return person, nil
}

//Helper Funcion
func NewPersonService(repository domain.PersonRepository) DefaultPersonService {
	return DefaultPersonService{Repo: repository}
}
