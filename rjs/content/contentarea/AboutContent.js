import AbstractContent from "./AbstractContent";

export default class AboutContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "about";
        this.contentTitle = "About Me"
    }
}