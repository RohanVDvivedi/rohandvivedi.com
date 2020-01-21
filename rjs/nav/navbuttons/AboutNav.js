import AbstractNav from "./AbstractNav";

export default class AboutNav extends AbstractNav {
    constructor(props) {
        super(props);
        this.name = "about"
        this.navTitle = "About"
    }
}