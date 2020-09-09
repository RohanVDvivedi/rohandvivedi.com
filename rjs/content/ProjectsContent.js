import React from "react";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";

class ProjectListerComponent extends React.Component {
    render() {
    	var hyperlinks = (this.props.project.Hyperlinks == null) ? [] : this.props.project.Hyperlinks
    	var thumbImage = hyperlinks.find((link) => {return link.Name.toLowerCase() == "thumbnail" && link.LinkType == "IMAGE"})
    	var GithubRepository = hyperlinks.find((link) => {return link.Name == this.props.project.Name && link.LinkType == "GITHUB"})

    	var GithubRepositories = hyperlinks.filter((link) => {return link.LinkType == "GITHUB"});
    	var YoutubeVideos = hyperlinks.filter((link) => {return link.LinkType == "YOUTUBE"});
        return (
            <div class="project-lister-element flex-col-container set_sub_content_background_color generic-content-box-border"
            		style={{
            			justifyContent: "space-between"
            		}}>
                <h1 class="project-lister-element-name">{this.props.project.Name}</h1>

                {(thumbImage != null) ? (<img class="project-lister-element-image" src={thumbImage.Href}/>) : ""}

	            <h3 class="project-lister-element-description">{this.props.project.Descr}</h3>

				{GithubRepositories != null && GithubRepositories.length > 0 ?
                	(<div class="flex-row-container" style={{justifyContent: "space-around",
                											alignItems: "center",}}>
                		{GithubRepositories.map((link) => {
                			return (<Icon path={link.Href} iconPath={"/icon/github.png"} 
                					infoBoxText={link.Name} height="35px" width="35px" padding="5px" />)
                		})}
                	</div>) : ""
            	}

				{YoutubeVideos != null && YoutubeVideos.length > 0 ?
					(<div class="flex-row-container" style={{justifyContent: "space-around",
                											alignItems: "center",}}>
                		{YoutubeVideos.map((link) => {
                			return (<Icon path={link.Href} iconPath={"/icon/youtube.png"} 
                					infoBoxText={link.Descr} height="35px" width="35px" padding="5px" />)
                		})}
                	</div>) : ""
				}
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
	onCategoryDrawerClicked() {
		this.updateState({
				searchTextBox: this.state.searchTextBox,
				showDropdownContent: !this.state.showDropdownContent,
				selectedProjectCategories: this.state.selectedProjectCategories,
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
	enterKeyClicked(event) {
		if (event.keyCode === 13) {event.preventDefault();
			this.searchButtonClicked();}
	}
	render() {
		// convert list of objects to list of strings
		this.categories = this.state.api_response_body.map(function(category){return category.Category});
		return (<div style={{display:"flex",justifyContent:"center"}}>
					<div class="search-container flex-row-container set_sub_content_background_color">
	                	<input class="search-text-selector" type="text" 
	                		placeholder="Search projects" value={this.state.searchTextBox}
	                		onChange={this.onSearchBoxTyping.bind(this)} 
	                		onKeyUp={this.enterKeyClicked.bind(this)}/>
						<div class={"search-categories-selector dropdown-container generic-content-box-hovering-emboss-border " + (this.state.showDropdownContent ? "show-dropdown-content" : "") }>
							<div class="flex-row-container" style={{justifyContent: "space-evenly", alignItems: "center"}} onClick={this.onCategoryDrawerClicked.bind(this)}>
								<div>Categories</div>
								<Icon path="#" iconPath="/icon/up-chevron.png" height="15px" width="15px" padding="2px" />
							</div>
							<div class="dropdown-content set_sub_content_background_color">
								<div id="select-all-categories" class={"search-category " + ((this.state.selectedProjectCategories.length==0) ? "active" : "")}
								 onClick={this.categoryClicked.bind(this, "")}>All</div>
								{this.categories.map((projectCategory) => {
									return (<div class={"search-category " + ((this.state.selectedProjectCategories.includes(projectCategory)) ? "active" : "")}
										onClick={this.categoryClicked.bind(this, projectCategory)}>{projectCategory}</div>);
								})}
							</div>
						</div>
	                	<div class="search-button generic-content-box-hovering-emboss-border" 
	                		onClick={this.searchButtonClicked.bind(this)}>Search</div>
	                </div>
                </div>);
	}
}

export default class ProjectsContent extends ApiComponent {
	constructor(props) {
		super(props)
		this.queryString = "get_all=true";
	}
	searchQueryStringBuiltCallback(queryString) {
		this.queryString = queryString;
		this.makeApiCallAndReRender()
	}
	apiPath() {
		const basePath = "/api/project";
		const queryParam = "get_hyperlinks=true";
        return [basePath, [this.queryString, queryParam].join("&")].join("?");
    }
    bodyDataBeforeApiFirstResponds() {
    	return [{Name: "Loading",Descr: "Loading Description",}];
    }
    render() {
    	var projects = this.state.api_response_body;
        return (
            <div class="content-root-container content-root-background">
                
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