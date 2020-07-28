import React from "react";

import Project from "./Project"

export default class ProjectsContent extends React.Component {
    render() {
        var projectNames = [
            "project-name","project-name","project-name","project-name","project-name","project-name",
            "project-name","project-name","project-name","project-name","project-name","project-name",
            "project-name","project-name","project-name","project-name","project-name","project-name",
        ];
        return (
            <div class="content-root-background">
                <div style={{height: "65px"}}></div>
                <div class="grid-container">
                        {projectNames.map(function(projectName, i){
                            return <Project projectName={projectName} />;
                        })}
                </div>
            </div>
        );
    }
}