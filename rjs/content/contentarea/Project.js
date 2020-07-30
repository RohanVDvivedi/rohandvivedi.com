import React from "react";
import ApiComponent from "../../utility/ApiComponent";
import Icon from "../../utility/Icon";

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
            <div class="project flex-col-container set_sub_content_background_color">
                <div id="project-name" style={{
                    textAlign: "center",
                    fontSize: "20px",
                    fontFamily: "Arial, Helvetica, sans-serif",
                    fontWeight: "700",
                }}>
                    {project.Name}
                </div>

                <div id="project-thumbnail" style={{
                    height: "45%",
                    width: "100%",
                    backgroundImage: "url('" + project.ProjectImage + "')",
                }}></div>
                
                <div id="project-description" style={{
                    fontFamily: "Arial, Helvetica, sans-serif",
                    textAlign: "center",
                }}>
                    {project.ProjectDescriptionShort}
                </div>

                <div class="flex-row-container"
                    style={{
                        justifyContent: "space-around",
                        alignItems: "baseline",
                    }}>
                    <a href={project.GithubLink} target="_blank" style={{display: "block",}}>
                        <Icon height="35px" width="35px" iconPath="/icon/github.png"/>
                    </a>
                    <a href={project.YoutubeLink} target="_blank" style={{display: "block",}}>
                        <Icon height="35px" width="35px" iconPath="/icon/youtube.png"/>
                    </a>
                </div>
            </div>
        );
    }
}