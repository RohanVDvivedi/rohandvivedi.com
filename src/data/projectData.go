package data

// Project model
type Project struct {
	Id NullInt64
	Name NullString
	Descr NullString
	GithubLink NullString
	YoutubeLink NullString
	ImageLink NullString
	ProjectOwner NullInt64
}

func projectSelectBaseQuery() string {
	return "select projects.id, projects.name, projects.descr, projects.github_link, projects.youtube_link, projects.image_link, projects.project_owner from projects ";
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
	return baseScanProject(Db.QueryRow(projectSelectBaseQuery() + "where projects.id = ?", id));
}

func GetProjectByName(name string) *Project {
	return baseScanProject(Db.QueryRow(projectSelectBaseQuery() + "where projects.name = ?", name));
}

func GetAllProjects() []Project {
	p := []Project{}
	rows, _ := Db.Query(projectSelectBaseQuery())
	defer rows.Close()
	for rows.Next() {
		project_p := baseScanProject(rows)
		if project_p != nil {
			p = append(p, *project_p)
		}
	}
	return p;
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
	Id NullInt64
	Href NullString
	Descr NullString
	ProjectId NullInt64
}

func projectHyperlinkSelectBaseQuery() string {
	return "select project_hyperlinks.id, project_hyperlinks.href, project_hyperlinks.descr, project_hyperlinks.project_id from project_hyperlinks ";
}

func baseScanProjectHyperlink(r Row) *ProjectHyperlink {
	ph := ProjectHyperlink{};
	err := r.Scan(&ph.Id, &ph.Href, &ph.Descr, &ph.ProjectId);
	if err != nil {
		return nil
	}
	return &ph;
}

func (p *Project) GetProjectHyperlinks() []ProjectHyperlink {
	ph := []ProjectHyperlink{}
	rows, _ := Db.Query(projectHyperlinkSelectBaseQuery() + "where project_hyperlinks.project_id = ?", p.Id)
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
	Id NullInt64
	Category NullString
	Descr NullString
}

func projectCategorySelectBaseQuery() string {
	return "select project_categories.id, project_categories.category_name, project_categories.descr from project_categories ";
}

func baseScanProjectCategory(r Row) *ProjectCategory {
	pc := ProjectCategory{};
	err := r.Scan(&pc.Id, &pc.Category, &pc.Descr);
	if err != nil {
		return nil
	}
	return &pc;
}

func GetProjectCategoryByName(name string) *ProjectCategory {
	return baseScanProjectCategory(Db.QueryRow(projectCategorySelectBaseQuery() + "where project_categories.category_name = ?", name));
}

func GetAllCategories() []ProjectCategory {
	pc := []ProjectCategory{}
	rows, _ := Db.Query(projectCategorySelectBaseQuery())
	defer rows.Close()
	for rows.Next() {
		projectcategory_p := baseScanProjectCategory(rows)
		if projectcategory_p != nil {
			pc = append(pc, *projectcategory_p)
		}
	}
	return pc;
}

func (p *Project) GetProjectCategories() []ProjectCategory {
	pc := []ProjectCategory{}
	rows, _ := Db.Query(projectCategorySelectBaseQuery() + 
	` join project_category_project on project_category_project.project_category_id = project_categories.id
	 join projects on projects.id = project_category_project.project_id where projects.id = ?`, p.Id)
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
	 join project_categories on project_categories.id = project_category_project.project_category_id where project_categories.id = ?`, pc.Id)
	defer rows.Close()
	for rows.Next() {
		project_p := baseScanProject(rows)
		if project_p != nil {
			p = append(p, *project_p)
		}
	}
	return p;
}

func GetProjectsForCategoryNames(categories_str []string) []Project {
	p := []Project{}

	categories := convertToInterfaceSlice(categories_str)

	rows, _ := Db.Query(projectSelectBaseQuery() + 
	` join project_category_project on project_category_project.project_id = projects.id
	 join project_categories on project_categories.id = project_category_project.project_category_id 
	 where project_categories.category_name in (` + getRepeatedQueryParamHolders(len(categories)) + ") group by projects.id", categories...)
	defer rows.Close()
	for rows.Next() {
		project_p := baseScanProject(rows)
		if project_p != nil {
			p = append(p, *project_p)
		}
	}
	return p;
}