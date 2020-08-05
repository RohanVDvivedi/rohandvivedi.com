package data

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

func baseScanPerson(r Row) *Person {
	p := Person{};
	err := r.Scan(&p.Id, &p.Fname, &p.Lname, &p.Email, &p.UserType);
	if err != nil {
		return nil
	}
	return &p;
}

func GetOwner() *Person {
	return baseScanPerson(Db.QueryRow(personSelectBaseQuery() + "where type = ?", "owner"));
}

func GetPersonById(id int) *Person {
	return baseScanPerson(Db.QueryRow(personSelectBaseQuery() + "where id = ?", id));
}

func GetPersonByName(fname string, lname string) *Person {
	return baseScanPerson(Db.QueryRow(personSelectBaseQuery() + "where fname = ? and lname = ?", fname, lname));
}

func UpdatePerson(p *Person) {
	Db.Exec("update persons set fname = ?, lname = ?, email = ?, type = ? where id = ?", p.Fname, p.Lname, p.Email, p.UserType, p.Id);
}

func InsertPerson(p *Person) {
	Db.Exec("insert into persons (fname, lname, email, type) values (?,?,?,?)", p.Fname, p.Lname, p.Email, p.UserType);
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

func baseScanSocial(r Row) *Social {
	s := Social{};
	err := r.Scan(&s.Id, &s.Descr, &s.ProfileLink, &s.LinkType, &s.PersonId);
	if err != nil {
		return nil
	}
	return &s;
}

func (p *Person) FindSocials() []Social {
	s := []Social{}
	rows, _ := Db.Query(socialSelectBaseQuery() + "where person_id = ?", p.Id)
	defer rows.Close()
	for rows.Next() {
		social_p := baseScanSocial(rows)
		if social_p != nil {
			s = append(s, *social_p)
		}
	}
	return s;
}