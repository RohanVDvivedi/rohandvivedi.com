import React from "react";
import AbstractContent from "./AbstractContent";

export default class SocialContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "social";
        this.contentTitle = "Social Links"
    }
    content() {
        return (
            <>
            </>
        );
    }
}