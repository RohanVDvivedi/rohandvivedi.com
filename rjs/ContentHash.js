import AboutContent from "./content/contentarea/AboutContent";
import ContactContent from "./content/contentarea/ContactContent";
import ProjectsContent from "./content/contentarea/ProjectsContent";
import PastContent from "./content/contentarea/PastContent";
import AboutThisAppContent from "./content/contentarea/AboutThisAppContent";
import AdminContent from "./content/contentarea/AdminContent";


var ContentHash = {
	/*
	just as an example, this is you you would add iconned page
	"app": {
		"text": "{;}",
		"nav_pos": "left",
		"component": AboutThisAppContent,
	},
	*/
	"about": {
		"text": "About",
		"nav_pos": "right",
		"component": AboutContent,
	},
	"projects": {
		"text": "Projects",
		"nav_pos": "right",
		"component": ProjectsContent,
	},
	"past": {
		"text": "Past",
		"nav_pos": "right",
		"component": PastContent,
	},
	"contact": {
		"text": "Contact",
		"nav_pos": "right",
		"component": ContactContent,
	},
	/*
	just as an example, this is you you would add iconned page
	"admin": {
		"icon": "/icon/spanner.svg",
		"nav_pos": "right",
		"component": AdminContent,
	},*/
};

export default ContentHash;