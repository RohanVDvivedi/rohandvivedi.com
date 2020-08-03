package data

import (
    "database/sql"
)

func InitializeSchema(db *sql.DB) {
    statement, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS persons (
									id INTEGER PRIMARY KEY AUTOINCREMENT, 
									fname VARCHAR(255) NOT NULL, 
									lname VARCHAR(255) NOT NULL, 
									email VARCHAR(128) NOT NULL, 
									ph_no VARCHAR(30), 
									type VARCHAR(30)
								)`);
    statement.Exec()

    statement, _ = db.Prepare(`CREATE TABLE IF NOT EXISTS social (
									id INTEGER PRIMARY KEY AUTOINCREMENT, 
									descr VARCHAR(512), 
									profile_link VARCHAR(512) NOT NULL, 
									link_type VARCHAR(128) NOT NULL, 
									person_id INTEGER,
									FOREIGN KEY(person_id) REFERENCES persons(id)
								)`);
    statement.Exec()

    statement, _ = db.Prepare(`CREATE TABLE IF NOT EXISTS projects (
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									name VARCHAR(128) NOT NULL,
									descr VARCHAR(512) NOT NULL,
									project_type VARCHAR(255) NOT NULL,
									github_link VARCHAR(512),
									youtube_link VARCHAR(512),
									project_owner,
									FOREIGN KEY(project_owner) REFERENCES persons(id)
								)`);
    statement.Exec()

    statement, _ = db.Prepare(`CREATE TABLE IF NOT EXISTS project_hyperlinks (
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									href VARCHAR(512) NOT NULL,
									descr VARCHAR(512) NOT NULL,
									project_id,
									FOREIGN KEY(project_id) REFERENCES project(id)
								)`);
    statement.Exec()

    statement, _ = db.Prepare(`CREATE TABLE IF NOT EXISTS person_project (
									person_id INTEGER,
									project_id INTEGER,
									FOREIGN KEY(person_id) REFERENCES person(id),
									FOREIGN KEY(project_id) REFERENCES project(id)
								)`);
    statement.Exec()
}

