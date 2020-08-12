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
type ProjectHyperlink struct {
	Id int
	Href string
	Descr string
	ProjectId int
}

func projectHyperlinkSelectBaseQuery() string {
	return "select id, href, descr, project_id from project_hyperlinks ";
}

func baseScanProjectHyperlink(r Row) *ProjectHyperlink {
	ph := ProjectHyperlink{};
	err := r.Scan(&ph.Id, &ph.Href, &ph.Descr, &ph.ProjectId);
	if err != nil {
		return nil
	}
	return &ph;
}

func (p *Project) GetHyperlinks() []ProjectHyperlink {
	ph := []ProjectHyperlink{}
	rows, _ := Db.Query(projectHyperlinkSelectBaseQuery() + "where project_id = ?", p.Id)
	defer rows.Close()
	for rows.Next() {
		projecthyperlink_p := baseScanProjectHyperlink(rows)
		if projecthyperlink_p != nil {
			ph = append(ph, *projecthyperlink_p)
		}
	}
	return ph;
}

// Categories
type ProjectCategory struct {
	Id int
	Category string
	Descr string
}

func projectCategorySelectBaseQuery() string {
	return "select id, category_name, descr from project_categories ";
}

func baseScanProjectCategory(r Row) *ProjectCategory {
	pc := ProjectCategory{};
	err := r.Scan(&pc.Id, &pc.Category, &pc.Descr);
	if err != nil {
		return nil
	}
	return &pc;
}

func (p *Project) GetProjectCategories() []ProjectCategory {
	pc := []ProjectCategory{}
	rows, _ := Db.Query(projectCategorySelectBaseQuery() + 
	` join project_category_project on project_category_project.project_category_id = project_category.id
	 join projects on projects.id = project_category_project.project_id where project.id = ?`, p.Id)
	defer rows.Close()
	for rows.Next() {
		projectcategory_p := baseScanProjectCategory(rows)
		if projectcategory_p != nil {
			pc = append(pc, *projectcategory_p)
		}
	}
	return pc;
}

/*select pr.name, cat.category_name from projects as pr join project_category_project as pc on (pc.project_id = pr.id) join project_categories as cat on (cat.id = pc.project_category_id) where cat.id = 2;*/

func (pc *ProjectCategory) GetProjects() []Project {
	p := []Project{}
	rows, _ := Db.Query(projectSelectBaseQuery() + 
	` join project_category_project on project_category_project.project_id = projects.id
	 join project_categories on project_categories.id = project_category_project.project_category_id where project_category.id = ?`, pc.Id)
	defer rows.Close()
	for rows.Next() {
		project_p := baseScanProject(rows)
		if project_p != nil {
			p = append(p, *project_p)
		}
	}
	return p;
}