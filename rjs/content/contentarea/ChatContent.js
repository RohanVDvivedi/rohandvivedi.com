import AbstractContent from "./AbstractContent";

export default class ChatContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "chat";
        this.contentTitle = "Chat Live"
    }
}