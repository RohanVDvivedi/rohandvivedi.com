package data

// Project model
type Project struct {
	Id int
	Name string
	Descr string
	GithubLink string
	YoutubeLink string
	ImageLink string
	ProjectOwner int
}

func projectSelectBaseQuery() string {
	return "select id, name, descr, github_link, youtube_link, image_link, project_owner from projects ";
}

func baseScanProject(r Row) *Project {
	p := Project{};
	err := r.Scan(&p.Id, &p.Name, &p.Descr, &p.GithubLink, &p.YoutubeLink, &p.ImageLink, &p.ProjectOwner);
	if err != nil {
		return nil
	}
	return &p;
}

func GetProjectById(id int) *Project {
	return baseScanProject(Db.QueryRow(projectSelectBaseQuery() + "where id = ?", id));
}

func GetProjectByName(name string) *Project {
	return baseScanProject(Db.QueryRow(projectSelectBaseQuery() + "where name = ?", name));
}

func UpdateProject(p *Project) {
	Db.Exec("update projects set name = ?, descr = ?, github_link = ?, youtube_link = ?, image_link = ?, project_owner = ? where id = ?",
		p.Name, p.Descr, p.GithubLink, p.YoutubeLink, p.ImageLink, p.ProjectOwner, p.Id);
}

func InsertProject(p *Project) {
	Db.Exec("insert into projects (name, descr, github_link, youtube_link, image_link, project_owner) values (?,?,?,?)",
		p.Name, p.Descr, p.GithubLink, p.YoutubeLink, p.ImageLink, p.ProjectOwner);
}

// Project's hyperlinks

// Categories

/*select pr.name, cat.category_name from projects as pr join project_category_project as pc on (pc.project_id = pr.id) join project_categories as cat on (cat.id = pc.project_category_id) where cat.id = 2;*/