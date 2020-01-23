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
        var projectNames = [
            "project-name","project-name","project-name","project-name","project-name","project-name",
            "project-name","project-name","project-name","project-name","project-name","project-name",
            "project-name","project-name","project-name","project-name","project-name","project-name",
        ];
        return (
            <div id={this.getContentId()} class="content-component">
                <div class="flex-row-container"
                style={{
                    flexWrap: "wrap",
                    justifyContent: "flex-start",
                    alignContent: "space-evenly",
                }}>
                    {projectNames.map(function(projectName, i){
                        return <Project projectName={projectName} />;
                    })}
                </div>
            </div>
        );
    }
}