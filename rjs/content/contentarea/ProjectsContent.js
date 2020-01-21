import AbstractContent from "./AbstractContent";

export default class ProjectsContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "projects";
        this.contentTitle = "My Projects"
    }
}