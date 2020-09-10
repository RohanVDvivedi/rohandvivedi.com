package data

import (
	"strings"
)

// Project model
type Project struct {
	Id NullInt64
	Name NullString
	Descr NullString
	ProgrLangs NullString
	LibsUsed NullString
	SkillSets NullString
	ProjectOwner NullInt64
}

func projectSelectBaseQuery() string {
	return "select projects.id, projects.name, projects.descr, projects.progr_langs, projects.libs_used, projects.skill_sets, projects.project_owner from projects ";
}

func baseScanProject(r Row) *Project {
	p := Project{};
	err := r.Scan(&p.Id, &p.Name, &p.Descr, &p.ProgrLangs, &p.LibsUsed, &p.SkillSets, &p.ProjectOwner);
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

func GetProjectsByNames(names_str []string) []Project {
	p := []Project{}
	names := convertToInterfaceSlice(names_str)
	rows, _ := Db.Query(projectSelectBaseQuery() + "where projects.name in (" + getRepeatedQueryParamHolders(len(names)) + ")", names...)
	defer rows.Close()
	for rows.Next() {
		project_p := baseScanProject(rows)
		if project_p != nil {
			p = append(p, *project_p)
		}
	}
	return p;
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
	Db.Exec("update projects set name = ?, descr = ?, progr_langs = ?, libs_used = ?, skill_sets = ?, project_owner = ? where id = ?", p.Name, p.Descr, p.ProgrLangs, p.LibsUsed, p.SkillSets, p.ProjectOwner, p.Id);
}

func InsertProject(p *Project) {
	_, err := Db.Exec("insert into projects (name, descr, progr_langs, libs_used, skill_sets, project_owner) values (?,?,?,?,?,?)", p.Name, p.Descr, p.ProgrLangs, p.LibsUsed, p.SkillSets, p.ProjectOwner);
	if(err == nil) {
		*p = *GetProjectByName(p.Name.NullString.String)
	}
}

// Project's hyperlinks
type ProjectHyperlink struct {
	Id NullInt64
	Name NullString
	Href NullString
	LinkType NullString
	Descr NullString
	ProjectId NullInt64
}

func projectHyperlinkSelectBaseQuery() string {
	return "select project_hyperlinks.id, project_hyperlinks.name, project_hyperlinks.href, project_hyperlinks.link_type, project_hyperlinks.descr, project_hyperlinks.project_id from project_hyperlinks ";
}

func baseScanProjectHyperlink(r Row) *ProjectHyperlink {
	ph := ProjectHyperlink{};
	err := r.Scan(&ph.Id, &ph.Name, &ph.Href, &ph.LinkType, &ph.Descr, &ph.ProjectId);
	if err != nil {
		return nil
	}
	return &ph;
}

func UpdateProjectHyperlink(p *ProjectHyperlink) {
	Db.Exec("update project_hyperlinks set name = ?, href = ?, link_type = ?, descr = ?, project_id = ? where id = ?", p.Name, p.Href, p.LinkType, p.Descr, p.ProjectId, p.Id);
}

func InsertProjectHyperlink(p *ProjectHyperlink) {
	_, err := Db.Exec("insert into project_hyperlinks (name, href, link_type, descr, project_id) values (?,?,?,?,?)", p.Name, p.Href, p.LinkType, p.Descr, p.ProjectId);
	if(err == nil) {
		*p = *baseScanProjectHyperlink(Db.QueryRow(projectHyperlinkSelectBaseQuery() + 
		"where project_hyperlinks.project_id = ? and project_hyperlinks.name = ? and project_hyperlinks.link_type = ?", p.ProjectId, p.Name, p.LinkType))
	}
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

func (p *Project) GetProjectGithubRepositoryLink() *ProjectHyperlink {
	githubRepoName := p.Name
	if(githubRepoName.Valid) {
		githubRepoName.String = strings.Replace(githubRepoName.String, " ", "-", -1)
	}
	return baseScanProjectHyperlink(Db.QueryRow(projectHyperlinkSelectBaseQuery() + 
	`where project_hyperlinks.project_id = ? and project_hyperlinks.name = ? and project_hyperlinks.link_type = ?`,
	 p.Id, githubRepoName, "GITHUB"))
}

func (p *Project) GetProjectGithubRepositoryLinks() []ProjectHyperlink {
	ph := []ProjectHyperlink{}
	rows, _ := Db.Query(projectHyperlinkSelectBaseQuery() + 
	"where project_hyperlinks.project_id = ? and project_hyperlinks.link_type = ?", p.Id, "GITHUB")
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