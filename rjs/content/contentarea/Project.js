import React from "react";
import ApiComponent from "../../utility/apicomponent/ApiComponent";

export default class Project extends ApiComponent {
    constructor(props) {
        super(props);
    }
    apiPath() {
        return "/api/project?projectName=" + this.props.projectName;
    }
    renderAfterApiSuccess() {
        var project = this.state.api_response_body;
        return (
            <div class="project flex-col-container">
                <div id="project-name">
                    {project.Name}
                </div>
                <div id="project-thumbnail" style={{
                    height: "45%",
                    width: "100%",
                    backgroundImage: "url('" + project.ProjectImage + "')",
                }}>
                </div>
                <div id="project-description">
                    {project.ProjectDescriptionShort}
                </div>
            </div>
        );
    }
}