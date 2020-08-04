package data

import (
    "database/sql"
)

type Person struct {
	Id int
	Fname string
	Lname string
	Email string
	PhNo string
	UserType string
}

type Social struct {
	Id int
	Descr string
	ProfileLink string
	LinkType string
	PersonId string
}

func personSelectBaseQuery() string {
	return "select id, fname, lname, email, type from persons ";
}

func baseScan(r *sql.Row) *Person {
	p := Person{};
	err := r.Scan(&p.Id, &p.Fname, &p.Lname, &p.Email, &p.UserType);
	if err != nil {
		return nil
	}
	return &p;
}

func GetOwner(db *sql.DB) *Person {
	return baseScan(db.QueryRow(personSelectBaseQuery() + "where type = ?", "owner"));
}

func GetPerson(id int, db *sql.DB) *Person {
	return baseScan(db.QueryRow(personSelectBaseQuery() + "where id = ?", id));
}

func UpdatePerson(p *Person, db *sql.DB) {
	db.Exec("update persons set fname = ?, lname = ?, email = ?, type = ? where id = ?", p.Fname, p.Lname, p.Email, p.UserType, p.Id);
}

func InsertPerson(p *Person, db *sql.DB) {
	db.Exec("insert into persons (fname, lname, email, type) values (?,?,?,?)", p.Fname, p.Lname, p.Email, p.UserType);
}