import AboutContent from "./content/contentarea/AboutContent";
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
		"drop_down": [
			"All",
			"Systems Programming (in Linux)",
			"Embedded System",
			"Robotics",
			"Computer Vision",
		]
	},
	"past": {
		"text": "Past",
		"nav_pos": "right",
		"component": PastContent,
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