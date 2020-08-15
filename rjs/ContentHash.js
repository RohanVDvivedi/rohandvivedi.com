import AboutContent from "./content/AboutContent";
import ProjectsContent from "./content/ProjectsContent";
import PastContent from "./content/PastContent";
import AboutThisAppContent from "./content/AboutThisAppContent";
import AdminContent from "./content/AdminContent";


var ContentHash = {
	"app": {
		"text": "{;}",
		"nav_pos": "left",
		"component": AboutThisAppContent,
		"route_path": "/pages/about_this_app",
		"pop_up_info": "About this app"
	},
	"about": {
		"text": "About",
		"nav_pos": "right",
		"component": AboutContent,
		"route_path": "/pages/about",
	},
	"projects": {
		"text": "Projects",
		"nav_pos": "right",
		"component": ProjectsContent,
		"route_path": "/pages/projects",
	},
	"past": {
		"text": "Past",
		"nav_pos": "right",
		"component": PastContent,
		"route_path": "/pages/past",
	},
	/*
	just as an example, this is how you would add iconned page
	"admin": {
		"icon": "/icon/spanner.svg",
		"nav_pos": "right",
		"component": AdminContent,
		"route_path": "/pages/admin",
	},*/
};

export default ContentHash;