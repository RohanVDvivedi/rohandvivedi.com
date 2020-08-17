package data

import "strings"

func SearchProjectsByQueryString(queryString string) []Project {
	p := []Project{}

	// cut strings on spaces
	queryStrings := strings.Fields(queryString)

	// convert each of them to wildcard %string%
	for i, s := range queryStrings {
		queryStrings[i] = "%" + s + "%"
	}

	// make two copies (one for name and one for description and one for category)
	queryStrings = append(queryStrings, queryStrings...)

	// builr parameters list
	queryParams := convertToInterfaceSlice(queryStrings)

	query := projectSelectBaseQuery() + " where "
	clause := getRepeatedwithDelimeter(" projects.name like ? ", "or", len(queryParams)/2) + 
		"or" + getRepeatedwithDelimeter(" projects.descr like ? ", "or", len(queryParams)/2)

	rows, _ := Db.Query(query + clause, queryParams...)
	defer rows.Close()
	for rows.Next() {
		project_p := baseScanProject(rows)
		if project_p != nil {
			p = append(p, *project_p)
		}
	}
	return p;
}