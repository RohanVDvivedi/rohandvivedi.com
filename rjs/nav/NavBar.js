import React from "react";

import GenericNavButton from "./navbuttons/GenericNavButton";

export default class NavBar extends React.Component {
    render() {
        return (
            <div id="nav-container" class="nav-container-style flex-row-container"
            style={{
                justifyContent: "flex-end",
            }}>
                <GenericNavButton name="about"      selected={this.props.selected} ifNavButtonClicked={this.props.ifNavButtonClicked}/>
                <GenericNavButton name="contact"    selected={this.props.selected} ifNavButtonClicked={this.props.ifNavButtonClicked}/>
                <GenericNavButton name="projects"   selected={this.props.selected} ifNavButtonClicked={this.props.ifNavButtonClicked}/>
                <GenericNavButton name="research"   selected={this.props.selected} ifNavButtonClicked={this.props.ifNavButtonClicked}/>
            </div>
        );
    }
}