import React from "react";

import ProjectSearchBar from "./ProjectSearchBar";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";

function removeAbbreviationInBrackets(str) {
    return (str.includes("(") && str.includes(")")) ? 
    (str.replace(str.slice(str.search("\\("), str.search("\\)") + 1), "").trim()) : str;
}

class ProjectListerComponent extends React.Component {
    render() {
    	var hyperlinks = (this.props.project.Hyperlinks == null) ? [] : this.props.project.Hyperlinks
    	var thumbImage = hyperlinks.find((link) => {return link.Name.toLowerCase() == "thumbnail" && link.LinkType == "IMAGE"})
    	var GithubRepository = hyperlinks.find((link) => {return link.Name == this.props.project.Name && link.LinkType == "GITHUB"})

    	var GithubRepositories = hyperlinks.filter((link) => {return link.LinkType == "GITHUB"});
    	var YoutubeVideos = hyperlinks.filter((link) => {return link.LinkType == "YOUTUBE"});
    	var ExternalLinks = hyperlinks.filter((link) => {return link.LinkType == "EXTERNAL_LINK"});

    	var categories = (this.props.project.Categories == null) ? [] : 
    	this.props.project.Categories.map(function(categ){return removeAbbreviationInBrackets(categ.Category)}).sort(function(a, b){return a.length - b.length;});

        return (
            <div class="project-lister-element flex-col-container set_sub_content_background_color generic-content-box-border"
            	style={{justifyContent: "space-between"}}>

                <h1 class="project-lister-element-name">{this.props.project.Name}</h1>

                {(thumbImage != null) ? (<img class="project-lister-element-image" src={thumbImage.Href}/>) : ""}

	            <h3 class="project-lister-element-description">{this.props.project.Descr}</h3>

	            {this.props.project.Categories != null && this.props.project.Categories.length > 0 ? 
		            (<div class="project-lister-element-categories">
		            	<span>Category:</span> {
		            		categories.map(function(categ){
		            			return (<span class="project-lister-element-category"> {categ} </span>)
		            		})
		            	}
	    	        </div>) : ""
	    	    }

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

				{ExternalLinks != null && ExternalLinks.length > 0 ?
					(<div class="flex-col-container" style={{justifyContent: "space-around",
															alignItems: "center",}}>
						{ExternalLinks.map((link) => {
							return (<a href={link.Href} target="_blank"> {link.Name} </a>)
						})}
					</div>) : ""
				}
            </div>
        );
    }
}

export default class ProjectsContent extends ApiComponent {
	constructor(props) {
		super(props)
		this.queryStringInit = "/api/project?get_hyperlinks=true&get_categories=true&get_all=true";
	}
	searchQueryStringBuiltCallback(queryString) {
		this.queryString = queryString;
		this.makeApiCallAndReRender()
	}
	apiPath() {
		if(this.queryString == null) {
			return this.queryStringInit;
		}
		return ["/api/search?get_hyperlinks=true&get_categories=true", this.queryString].join("&");
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
                	{(projects != null && projects.length > 0) ?
                        projects.map(function(project, i){return (<ProjectListerComponent project={project} />);}) : 
                        "Sorry, your search query returned no results..."}
                </div>
            </div>
        );
    }
}