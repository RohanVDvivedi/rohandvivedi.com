import React from "react";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";

class ProjectListerComponent extends ApiComponent {
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
            <div class="project-lister-element flex-col-container set_sub_content_background_color generic-content-box-border generic-content-box-hovering-emboss-border">
                <div class="project-lister-element-name">{project.Name}</div>
                <img class="project-lister-element-image" src={project.ImageLink}/>
	            <div class="project-lister-element-description">{project.Descr}</div>

                <div class="flex-row-container" style={{justifyContent: "space-around",
                										alignItems: "center",}}>
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
        var project_category = ["Systems programming (in linux)", "Embedded systems",
        "Robotics", "Databases", "Computer architecture"];
        return (
            <div class="content-container content-root-background">
                <div class="behind-nav"></div>
                
                <div class="flex-row-container" style={{
                	justifyContent: "center",
    				alignItems: "center",
    				fontSize: "20px",
    				padding: "8px",
                }}>
                	<input type="text" id="myText" placeholder="keywords to search projects" style={{font:"inherit"}} />
                	<div>
                		Select Category
                	</div>
                	<div>
                		Search
                	</div>
                </div>


                <div class="grid-container project-lister-contaier">
                        {projectNames.map(function(projectName, i){
                            return <ProjectListerComponent projectName={projectName} />;
                        })}
                </div>
            </div>
        );
    }
}