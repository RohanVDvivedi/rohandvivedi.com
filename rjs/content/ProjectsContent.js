import React from "react";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";

class ProjectListerComponent extends React.Component {
    render() {
        return (
            <div class="project-lister-element flex-col-container set_sub_content_background_color generic-content-box-border">
                <div class="project-lister-element-name">{this.props.project.Name}</div>
                <img class="project-lister-element-image" src={this.props.project.ImageLink}/>
	            <div class="project-lister-element-description">{this.props.project.Descr}</div>

                <div class="flex-row-container" style={{justifyContent: "space-around",
                										alignItems: "center",}}>
                    <Icon path={this.props.project.GithubLink} iconPath="/icon/github.png" infoBoxText={(<div>Open Github<br/>repository</div>)} height="35px" width="35px" padding="5px" />
                    <Icon path={this.props.project.YoutubeLink} iconPath="/icon/youtube.png" height="35px" width="35px" padding="5px" />
                </div>
            </div>
        );
    }
}

class ProjectSearchBar extends ApiComponent {
	constructor(props) {
		super(props)
		this.state =  Object.assign({} ,this.state, {
			searchTextBox: "",
			showDropdownContent: false,
			selectedProjectCategories: [],
		})
	}
	apiPath() {
        return "/api/all_categories";
    }
    bodyDataBeforeApiFirstResponds() {
    	return [];
    }
	categoryClicked(category) {
		var selects = [];
		if(category == "") {
			selects = [];
		} else {
			selects = Array.from(this.state.selectedProjectCategories);
			if(selects.includes(category)) {
				selects = selects.filter((cate) => {return cate != category});
			} else {
				selects.push(category);
				if(selects.length == this.categories.length){
					selects = []
				}
			}
		}
		this.updateState({
				searchTextBox: this.state.searchTextBox,
				showDropdownContent: selects.length > 0,
				selectedProjectCategories: selects,
			});
	}
	onSearchBoxTyping(e) {
		this.updateState({
				searchTextBox: e.target.value,
				showDropdownContent: this.state.showDropdownContent,
				selectedProjectCategories: this.state.selectedProjectCategories,
			});
	}
	generateSearchQueryString() {
		var searchQuery = new Array();

		var searchText = this.state.searchTextBox;
		if(searchText != null && searchText != ""){
			searchQuery[0] = "query=" + searchText;
		}

		var categories = ((this.state.selectedProjectCategories.length == 0) ? this.categories : this.state.selectedProjectCategories);
		if(categories != null && categories != ""){
			searchQuery[1] = "categories=" + categories.join(",");
		}

		// if search is pressed close the drop down
		this.updateState({
				searchTextBox: this.state.searchTextBox,
				showDropdownContent: false,
				selectedProjectCategories: this.state.selectedProjectCategories,
			});

		return searchQuery.filter((q) => {return q != null && q != ""}).join("&");
	}
	searchButtonClicked() {
		var queryString = this.generateSearchQueryString();
		this.props.searchQueryStringBuiltCallback(queryString);
	}
	render() {
		// convert list of objects to list of strings
		this.categories = this.state.api_response_body.map(function(category){return category.Category});
		return (<div style={{display:"flex",justifyContent:"center"}}>
					<div class="search-container flex-row-container set_sub_content_background_color">
	                	<input class="search-text-selector" type="text" placeholder="Search projects" onChange={this.onSearchBoxTyping.bind(this)} value={this.state.searchTextBox}/>
						<div class={"search-categories-selector dropdown-container generic-content-box-hovering-emboss-border " + (this.state.showDropdownContent ? "show-dropdown-content" : "") }>
							<div> Categories </div>
							<div class="dropdown-content set_sub_content_background_color">
								<div id="select-all-categories" class={"search-category " + ((this.state.selectedProjectCategories.length==0) ? "active" : "")}
								 onClick={this.categoryClicked.bind(this, "")}>All</div>
								{this.categories.map((projectCategory) => {
									return (<div class={"search-category " + ((this.state.selectedProjectCategories.includes(projectCategory)) ? "active" : "")}
										onClick={this.categoryClicked.bind(this, projectCategory)}>{projectCategory}</div>);
								})}
							</div>
						</div>
	                	<div class="search-button generic-content-box-hovering-emboss-border" onClick={this.searchButtonClicked.bind(this)}>Search</div>
	                </div>
                </div>);
	}
}

export default class ProjectsContent extends ApiComponent {
	constructor(props) {
		super(props)
		this.queryString = ((this.props.queryString==null)?"":this.props.queryString);
	}
	searchQueryStringBuiltCallback(queryString) {
		this.queryString = queryString;
		this.makeApiCallAndReRender()
	}
	apiPath() {
		const basePath = "/api/project";
        return [basePath, this.queryString].join("?");
    }
    bodyDataBeforeApiFirstResponds() {
    	return [{Name: "Loading",
				Descr: "Loading Description",
				GithubLink: "",
				YoutubeLink: "",
				ImageLink: "/img/pcb.jpeg",}];
    }
    render() {
    	var projects = this.state.api_response_body;
        return (
            <div class="content-container content-root-background">
                <div class="behind-nav"></div>
                
                <ProjectSearchBar searchQueryStringBuiltCallback={this.searchQueryStringBuiltCallback.bind(this)}/>

                <div class="grid-container project-lister-contaier">
                        {projects.map(function(project, i){
                            return <ProjectListerComponent project={project} />;
                        })}
                </div>
            </div>
        );
    }
}