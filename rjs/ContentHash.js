import AboutContent from "./content/AboutContent";
import ProjectsContent from "./content/ProjectsContent";
import PastContent from "./content/PastContent";
import AboutThisAppContent from "./content/AboutThisAppContent";
import AdminContent from "./content/AdminContent";


var ContentHash = {
	/*
	just as an example, this is you you would add iconned page
	"app": {
		"text": "{;}",
		"nav_pos": "left",
		"component": AboutThisAppContent,
		"route_path": "/about_this_app",
	},
	*/
	"about": {
		"text": "About",
		"nav_pos": "right",
		"component": AboutContent,
		"route_path": "/about",
	},
	"projects": {
		"text": "Projects",
		"nav_pos": "right",
		"component": ProjectsContent,
		"route_path": "/projects",
	},
	"past": {
		"text": "Past",
		"nav_pos": "right",
		"component": PastContent,
		"route_path": "/past",
	},
	/*
	just as an example, this is you you would add iconned page
	"admin": {
		"icon": "/icon/spanner.svg",
		"nav_pos": "right",
		"component": AdminContent,
		"route_path": "/admin",
	},*/
};

export default ContentHash;