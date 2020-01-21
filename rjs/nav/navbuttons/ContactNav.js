import React from "react";

import AbstractNav from "./AbstractNav";

export default class ContactNav extends AbstractNav {
    constructor(props) {
        super(props);
        this.name = "contact"
        this.navTitle = "Contact"
    }
}