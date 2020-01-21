import React from "react";

export default class AbstractContent extends React.Component {
    constructor(props) {
        super(props);
    }
    getNavId() {
        return this.name + "-nav"
    }
    getContentId() {
        return this.name + "-content"
    }
    render() {
        return (
            <div id={this.getContentId()} class="content-component">
                {this.content()}
            </div>
        );
    }
}