import React from "react";
import AbstractContent from "./AbstractContent";

export default class AboutContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "about";
        this.contentTitle = "About Me"
    }
    content() {
        return (
            <p>
                    Hi,
                <br />
                    I am Rohan
                <br />
            </p>
        );
    }
}