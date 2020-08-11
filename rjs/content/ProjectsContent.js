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
            <div class="project-lister-element flex-col-container set_sub_content_background_color generic-content-box-border">
                <div class="project-lister-element-name">{project.Name}</div>
                <img class="project-lister-element-image" src={project.ImageLink}/>
	            <div class="project-lister-element-description">{project.Descr}</div>

                <div class="flex-row-container" style={{justifyContent: "space-around",
                										alignItems: "center",}}>
                    <Icon path={project.GithubLink} iconPath="/icon/github.png" infoBoxText={(<div>Open Github<br/>repository</div>)} height="35px" width="35px" padding="5px" />
                    <Icon path={project.YoutubeLink} iconPath="/icon/youtube.png" height="35px" width="35px" padding="5px" />
                </div>
            </div>
        );
    }
}

class ProjectSearchBar extends React.Component {
	getSearchText() {
		return this.refs.searchTextBox.value;
	}
	getCategoriesSelected() {
		return "abc";
	}
	generateSearchQueryString() {
		var searchQuery = new Array();

		var searchText = this.getSearchText();
		if(searchText != null && searchText != ""){
			searchQuery[0] = "query=" + searchText;
		}

		var categories = this.getCategoriesSelected();
		if(categories != null && categories != ""){
			searchQuery[1] = "categories=" + categories;
		}

		return searchQuery.filter((q) => {return q != null && q != ""}).join("&");
	}
	searchButtonClicked() {
		var queryString = this.generateSearchQueryString();
		console.log(queryString);
	}
	render() {
		var projectCategories = ["Systems programming (in linux)", "Embedded systems","Robotics", "Databases", "Computer architecture"];

		return (<div style={{display:"flex",justifyContent:"center"}}>
					<div class="search-container flex-row-container set_sub_content_background_color">
	                	<input class="search-text-selector" type="text" placeholder="technical keywords" ref="searchTextBox"/>
						<div class="search-categories-selector dropdown-container generic-content-box-hovering-emboss-border">
							<div> Categories </div>
							<div class="dropdown-content set_sub_content_background_color" ref="parentOfCategories">
								<div id="select-all-categories" class="search-category active">All</div>
								{projectCategories.map((projectCategory) => {
									return (<div class="search-category">{projectCategory}</div>);
								})}
							</div>
						</div>
	                	<div class="search-button generic-content-box-hovering-emboss-border" onClick={this.searchButtonClicked.bind(this)}>Search</div>
	                </div>
                </div>);
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
                <div class="behind-nav"></div>
                
                <ProjectSearchBar />

                <div class="grid-container project-lister-contaier">
                        {projectNames.map(function(projectName, i){
                            return <ProjectListerComponent projectName={projectName} />;
                        })}
                </div>
            </div>
        );
    }
}