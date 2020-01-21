import React from "react";

export default class AboutContent extends React.Component {
    constructor(props) {
        super(props);
        this.name = "about";
    }
    render() {
        return (
            <div id={this.name + "-content"} class="content-component">
                About
            </div>
        );
    }
}