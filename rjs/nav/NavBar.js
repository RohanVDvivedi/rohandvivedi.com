import React from "react";

import ContentHash from "../ContentHash";

import GenericNavButton from "./navbuttons/GenericNavButton";

export default class NavBar extends React.Component {
    render() {
    	var leftButtons = Object.keys(ContentHash).filter((buttonName) => {return ContentHash[buttonName]["nav_pos"] == "left"})
    	var rightButtons = Object.keys(ContentHash).filter((buttonName) => {return ContentHash[buttonName]["nav_pos"] == "right"})
        return (
            <div id="nav-container" class="flex-row-container">

            	<div id="nav-container-left-buttons" class="flex-row-container">
	            	{leftButtons.map((buttonName) => {     
	           			return (<GenericNavButton
	           					name = {buttonName}
	           					buttonDescription={ContentHash[buttonName]} 
	           					isSelected={this.props.selected == buttonName}
	           					ifNavButtonClicked={this.props.ifNavButtonClicked}/>); 
	        			})
	            	}
	        	</div>

	        	<div id="nav-container-right-buttons" class="flex-row-container">
	            	{rightButtons.map((buttonName) => {     
	           			return (<GenericNavButton
	           					name = {buttonName}
	           					buttonDescription={ContentHash[buttonName]}
	           					isSelected={this.props.selected == buttonName}
	           					ifNavButtonClicked={this.props.ifNavButtonClicked}/>); 
	        			})
	            	}
	        	</div>

            </div>
        );
    }
}