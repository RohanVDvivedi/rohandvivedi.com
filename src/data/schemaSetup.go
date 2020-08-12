package data

import (
    "io/ioutil"
    "encoding/json"
)

func InitializeSchema() {

	// PARENT table persons
    statement, _ := Db.Prepare(`CREATE TABLE IF NOT EXISTS persons (
									id INTEGER PRIMARY KEY AUTOINCREMENT, 
									fname VARCHAR(255) NOT NULL, 
									lname VARCHAR(255) NOT NULL, 
									email VARCHAR(128) NOT NULL, 
									ph_no VARCHAR(30), 
									type VARCHAR(30),
									CONSTRAINT unique_person_name UNIQUE (fname, lname)
								)`);
    statement.Exec()

    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS socials (
									id INTEGER PRIMARY KEY AUTOINCREMENT, 
									descr VARCHAR(512), 
									profile_link VARCHAR(512) NOT NULL, 
									link_type VARCHAR(128) NOT NULL, 
									person_id INTEGER,
									FOREIGN KEY(person_id) REFERENCES persons(id)
								)`);
    statement.Exec()

    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS pasts (
									id INTEGER PRIMARY KEY AUTOINCREMENT, 
									organization VARCHAR(512) NOT NULL,
									organization_link VARCHAR(512),
									team_or_research_title VARCHAR(128),
									descr VARCHAR(512),
									research_paper_link  VARCHAR(512),
									from_date DATE, 
									to_date DATE, 
									person_id INTEGER,
									FOREIGN KEY(person_id) REFERENCES persons(id)
								)`);
    statement.Exec()

    // PARENT table projects
    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS projects (
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									name VARCHAR(128) NOT NULL,
									descr VARCHAR(512) NOT NULL,
									github_link VARCHAR(512),
									youtube_link VARCHAR(512),
									image_link VARCHAR(512),
									project_owner,
									FOREIGN KEY(project_owner) REFERENCES persons(id),
									CONSTRAINT unique_project_name UNIQUE (name)
								)`);
    statement.Exec()

    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS project_hyperlinks (
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									href VARCHAR(512) NOT NULL,
									descr VARCHAR(512) NOT NULL,
									project_id,
									FOREIGN KEY(project_id) REFERENCES projects(id)
								)`);
    statement.Exec()

    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS project_categories (
    								id INTEGER PRIMARY KEY AUTOINCREMENT,
    								category_name VARCHAR(128) NOT NULL,
    								descr VARCHAR(512),
    								CONSTRAINT unique_category_name UNIQUE (category_name)
    							)`);
    statement.Exec()

    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS project_category_project (
									project_category_id INTEGER,
									project_id INTEGER,
									FOREIGN KEY(project_category_id) REFERENCES project_categories(id),
									FOREIGN KEY(project_id) REFERENCES projects(id)
								)`);
    statement.Exec()

    statement, _ = Db.Prepare(`CREATE TABLE IF NOT EXISTS person_project (
									person_id INTEGER,
									project_id INTEGER,
									FOREIGN KEY(person_id) REFERENCES persons(id),
									FOREIGN KEY(project_id) REFERENCES projects(id)
								)`);
    statement.Exec()

    // below piece of code is required to update owner information
    // as and when needed, using the owner.json file
    // this is for convinience

    p_new_owner := Person{}
    data, _ := ioutil.ReadFile("./owner.json")
	_ = json.Unmarshal(data, &p_new_owner);
	p_new_owner.UserType = "owner"

	p := GetOwner();
    if(p != nil) {	// there exists an owner, just update everything except name
    	p_new_owner.Id = p.Id
    	if(p_new_owner.Fname != p.Fname || 
    		p_new_owner.Lname != p.Lname ||
    		p_new_owner.Email != p.Email ||
			p_new_owner.PhNo != p.PhNo) {	// update only if the fields change
	    	UpdatePerson(&p_new_owner)
	    }
    } else {	// insert an owner from the owner.json file
    	InsertPerson(&p_new_owner)
    }
}

