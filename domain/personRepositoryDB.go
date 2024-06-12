package domain

import (
	"database/sql"
	"log"
)

type DefaultPersonRepositoryDB struct {
	Client *sql.DB
}

func (d DefaultPersonRepositoryDB) FindByID(id int) (*Person, error) {
	personSQL := "SELECT p.name, ph.number AS phone_number, a.city, a.state, a.street1, a.street2, a.zip_code FROM person p JOIN phone ph ON p.id = ph.person_id JOIN  address_join aj ON p.id = aj.person_id JOIN address a ON aj.address_id = a.id WHERE  p.id=?';"
	var p Person
	row := d.Client.QueryRow(personSQL, id)
	err := row.Scan(&p.Name, &p.PhoneNumber, &p.City, &p.State, &p.Street1, &p.Street2, &p.ZipCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Println("Error while scaaning " + err.Error())
			return nil, err
		}
	}
	return &p, nil
}

//Helper Func
func NewPersonRepository(db *sql.DB) DefaultPersonRepositoryDB {
	return DefaultPersonRepositoryDB{Client: db}
}
