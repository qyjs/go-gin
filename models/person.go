package models

import (
	basic "gin/databases"

	"github.com/google/uuid"
)

// Person ...
type Person struct {
	Id        string `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

// AddPerson ...
func (p *Person) AddPerson() (id int64, err error) {
	p.Id = uuid.New().String()
	rs, err := basic.SQLDB.Exec("INSERT INTO person(first_name, last_name, id) VALUES (?, ?, ?)", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	rs, err := basic.SQLDB.Exec("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}

	ra, err = rs.RowsAffected()
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := basic.SQLDB.Exec("DELETE FROM person WHERE id =", p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

// GetPersons ...
func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := basic.SQLDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

// GetPerson ...
func (p *Person) GetPerson() (err error) {
	basic.SQLDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", p.Id).Scan(
		&p.Id,
		&p.FirstName,
		&p.LastName,
	)
	return
}
