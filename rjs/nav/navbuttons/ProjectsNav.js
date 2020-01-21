import AbstractNav from "./AbstractNav";

export default class ProjectsNav extends AbstractNav {
    constructor(props) {
        super(props);
        this.name = "projects"
        this.navTitle = "Projects"
    }
}