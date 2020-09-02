package data

// Person model
type Person struct {
	Id NullInt64
	Fname NullString
	Lname NullString
	Email NullString
	PhNo NullString
	UserType NullString
}

func personSelectBaseQuery() string {
	return "select persons.id, persons.fname, persons.lname, persons.email, persons.type from persons ";
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
	return baseScanPerson(Db.QueryRow(personSelectBaseQuery() + "where persons.type = ?", "owner"));
}

func GetPersonById(id int) *Person {
	return baseScanPerson(Db.QueryRow(personSelectBaseQuery() + "where persons.id = ?", id));
}

func GetPersonByName(fname string, lname string) *Person {
	return baseScanPerson(Db.QueryRow(personSelectBaseQuery() + "where persons.fname = ? and persons.lname = ?", fname, lname));
}

func UpdatePerson(p *Person) {
	Db.Exec("update persons set fname = ?, lname = ?, email = ?, type = ? where id = ?", p.Fname, p.Lname, p.Email, p.UserType, p.Id);
}

func InsertPerson(p *Person) {
	_, err := Db.Exec("insert into persons (fname, lname, email, type) values (?,?,?,?)", p.Fname, p.Lname, p.Email, p.UserType);
	if(err == nil) {
		*p = *GetPersonByName(p.Name)
	}
}

// Person's Social-s
type Social struct {
	Id NullInt64
	Descr NullString
	Username NullString
	ProfileLink NullString
	LinkType NullString
	PersonId NullInt64
}

func socialSelectBaseQuery() string {
	return "select socials.id, socials.descr, socials.username, socials.profile_link, socials.link_type, socials.person_id from socials ";
}

func baseScanSocial(r Row) *Social {
	s := Social{};
	err := r.Scan(&s.Id, &s.Descr, &s.Username, &s.ProfileLink, &s.LinkType, &s.PersonId);
	if err != nil {
		return nil
	}
	return &s;
}

func (p *Person) FindSocials() []Social {
	s := []Social{}
	rows, _ := Db.Query(socialSelectBaseQuery() + "where socials.person_id = ?", p.Id)
	defer rows.Close()
	for rows.Next() {
		social_p := baseScanSocial(rows)
		if social_p != nil {
			s = append(s, *social_p)
		}
	}
	return s;
}

// Person's Past-s
type Past struct {
	Id NullInt64
	Organization NullString
	OrganizationLink NullString
	PastType NullString
	Position NullString
	Team_or_ResearchTitle NullString
	Descr NullString
	ResearchPaperLink NullString
	FromDate NullString
	ToDate NullString
	PersonId NullInt64
}

func pastSelectBaseQuery() string {
	return "select pasts.id, pasts.organization, pasts.organization_link, pasts.past_type, pasts.position, pasts.team_or_research_title, pasts.descr, pasts.research_paper_link, pasts.from_date, pasts.to_date, pasts.person_id from pasts ";
}

func baseScanPast(r Row) *Past {
	pst := Past{};
	err := r.Scan(&pst.Id, &pst.Organization, &pst.OrganizationLink, &pst.PastType, &pst.Position, &pst.Team_or_ResearchTitle, &pst.Descr, &pst.ResearchPaperLink, &pst.FromDate, &pst.ToDate, &pst.PersonId);
	if err != nil {
		return nil
	}
	return &pst;
}

func (p *Person) FindPasts() []Past {
	pst := []Past{}
	rows, _ := Db.Query(pastSelectBaseQuery() + "where pasts.person_id = ?", p.Id)
	defer rows.Close()
	for rows.Next() {
		past_p := baseScanPast(rows)
		if past_p != nil {
			pst = append(pst, *past_p)
		}
	}
	return pst;
}