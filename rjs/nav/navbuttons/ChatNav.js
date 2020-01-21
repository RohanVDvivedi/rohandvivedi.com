import React from "react";

import AbstractNav from "./AbstractNav";

export default class ChatNav extends AbstractNav {
    constructor(props) {
        super(props);
        this.name = "chat"
        this.navTitle = "Chat Live"
    }
}