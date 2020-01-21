import React from "react";

export default class AbstractContent extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div id={this.name + "-content"} class="content-component">
                {this.contentTitle}
            </div>
        );
    }
}