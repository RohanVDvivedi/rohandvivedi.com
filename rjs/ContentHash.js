import AboutContent from "./content/contentarea/AboutContent";
import ContactContent from "./content/contentarea/ContactContent";
import ProjectsContent from "./content/contentarea/ProjectsContent";
import PastContent from "./content/contentarea/PastContent";
import AboutThisAppContent from "./content/contentarea/AboutThisAppContent";
import AdminContent from "./content/contentarea/AdminContent";


var ContentHash = {
	"{ ; }": AboutThisAppContent,
	"About": AboutContent,
	"Projects": ProjectsContent,
	"Past": PastContent,
	"Contact": ContactContent,
	"Admin": AdminContent,
};

export default ContentHash;