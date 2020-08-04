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

func personSelectBaseQuery() string {
	return "select id, fname, lname, email, type from persons ";
}

func baseScanPerson(r *sql.Row) *Person {
	p := Person{};
	err := r.Scan(&p.Id, &p.Fname, &p.Lname, &p.Email, &p.UserType);
	if err != nil {
		return nil
	}
	return &p;
}

func GetOwner(db *sql.DB) *Person {
	return baseScanPerson(db.QueryRow(personSelectBaseQuery() + "where type = ?", "owner"));
}

func GetPerson(id int, db *sql.DB) *Person {
	return baseScanPerson(db.QueryRow(personSelectBaseQuery() + "where id = ?", id));
}

func UpdatePerson(p *Person, db *sql.DB) {
	db.Exec("update persons set fname = ?, lname = ?, email = ?, type = ? where id = ?", p.Fname, p.Lname, p.Email, p.UserType, p.Id);
}

func InsertPerson(p *Person, db *sql.DB) {
	db.Exec("insert into persons (fname, lname, email, type) values (?,?,?,?)", p.Fname, p.Lname, p.Email, p.UserType);
}

type Social struct {
	Id int
	Descr string
	ProfileLink string
	LinkType string
	PersonId int
}

func socialSelectBaseQuery() string {
	return "select id, descr, profile_link, link_type, person_id from social ";
}

func baseScanSocial(r *sql.Rows) *Social {
	s := Social{};
	err := r.Scan(&s.Id, &s.Descr, &s.ProfileLink, &s.LinkType, &s.PersonId);
	if err != nil {
		return nil
	}
	return &s;
}

func (p *Person) FindSocials(db *sql.DB) []Social {
	s := []Social{}
	rows, _ := db.Query(socialSelectBaseQuery() + "where person_id = ?", p.Id)
	defer rows.Close()
	for rows.Next() {
		social_p := baseScanSocial(rows)
		if social_p != nil {
			s = append(s, *social_p)
		}
	}
	return s;
}