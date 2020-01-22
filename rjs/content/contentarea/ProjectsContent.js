import React from "react";
import AbstractContent from "./AbstractContent";

import Project from "./Project"

export default class ProjectsContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "projects";
        this.contentTitle = "My Projects"
    }
    render() {
        return (
            <div id={this.getContentId()} class="content-component flex-row-container"
            style={{
                paddingTop: "30px",
                flexWrap: "wrap",
                justifyContent: "space-around",
                alignContent: "space-around",
            }}>
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
                <Project projectName="project-name" />
            </div>
        );
    }
}