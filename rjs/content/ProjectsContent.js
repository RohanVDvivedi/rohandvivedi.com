import React from "react";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";

class Project extends ApiComponent {
    apiPath() {
        return "/api/project?name=" + this.props.projectName;
    }
    dataWhileApiResponds() {
    	return {
    		Name: "Loading",
    		Descr: "Loading Description",
    		GithubLink: "",
    		YoutubeLink: "",
    		ImageLink: "/img/pcb.jpeg",
    	};
    }
    render() {
        var project = this.state.api_response_body;
        return (
            <div class="flex-col-container set_sub_content_background_color generic-content-box-border generic-content-box-hovering-emboss-border"
            	style={{
            		fontFamily: "Capriola, Helvetica, sans-serif",
            		margin: "4%",
    				padding: "4%",
    				color: "rgb(50, 50, 50)",
            	}}
            	>
                <div id="project-name" style={{
                    textAlign: "center",
                    fontSize: "20px",
                    fontWeight: "700",
                    padding: "4%",
                }}>
                    {project.Name}
                </div>

                <img src={project.ImageLink} style={{padding: "4%", width: "92%"}}/>
	                
	            <div id="project-description" style={{
	                textAlign: "center",
	            	fontFamily: "lato, sans-serif",
	            	padding: "4%",
	            }}>
	                {project.Descr}
	            </div>

                <div class="flex-row-container"
                    style={{
                        justifyContent: "space-around",
                        alignItems: "center",
                    }}>
                    <Icon path={project.GithubLink} iconPath="/icon/github.png" height="35px" width="35px" padding="5px" />
                    <Icon path={project.YoutubeLink} iconPath="/icon/youtube.png" height="35px" width="35px" padding="5px" />
                </div>
            </div>
        );
    }
}

export default class ProjectsContent extends React.Component {
    render() {
        var projectNames = [
            "project-name","project-name","project-name","project-name","project-name","project-name",
            "project-name","project-name","project-name","project-name","project-name","project-name",
            "project-name","project-name","project-name","project-name","project-name","project-name",
        ];
        return (
            <div class="content-container content-root-background">
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