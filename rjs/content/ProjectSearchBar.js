import React from "react";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";

export default class ProjectSearchBar extends ApiComponent {
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
			searchQuery.push("query=" + searchText);
		}

		var categories = this.state.selectedProjectCategories;
		if(categories != null && categories.length > 0){
			searchQuery.push("categories=" + categories.join(","));
		}

		return (searchQuery.length == 0) ? null : searchQuery.join("&");
	}
	searchButtonClicked() {
		var queryString = this.generateSearchQueryString();

		// if search is pressed close the drop down, and remove all selection categories
		this.updateState({
				searchTextBox: this.state.searchTextBox,
				showDropdownContent: false,
				selectedProjectCategories: [],
			});
		
		this.props.searchQueryStringBuiltCallback(queryString);
	}
	enterKeyClicked(event) {
		if (event.keyCode === 13) {event.preventDefault();
			this.searchButtonClicked();}
	}
	render() {
		// convert list of objects to list of strings
		this.categories = this.state.api_response_body.map(function(category){return category.Category});
		return (<div class="loading-able-container" style={{display:"flex",justifyContent:"center"}}>

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
		               		onClick={this.searchButtonClicked.bind(this)}>
		               		Search
		               	</div>
		               	<div class={"search-loading" + ((this.props.loading)?"":"-hidden")}>
		               		<Icon path="#" iconPath="/icon/loading.gif" height="50px" width="50px" padding="0px" />
		            	</div>
		            </div>
                </div>);
	}
}