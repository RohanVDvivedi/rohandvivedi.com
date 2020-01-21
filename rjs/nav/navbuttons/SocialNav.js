import React from "react";

import AbstractNav from "./AbstractNav";

export default class SocialNav extends AbstractNav {
    constructor(props) {
        super(props);
        this.name = "social"
        this.navTitle = "Social"
    }
}