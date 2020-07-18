import React from "react";

import ContentHash from "../ContentHash";

import GenericNavButton from "./navbuttons/GenericNavButton";

export default class NavBar extends React.Component {
    render() {
        return (
            <div id="nav-container" class="nav-container-style flex-row-container"
            style={{
                justifyContent: "flex-end",
            }}>
            	{Object.keys(ContentHash).map((buttonName) => {     
           			return (
           				<GenericNavButton name={buttonName}
           					selected={this.props.selected}
           					ifNavButtonClicked={this.props.ifNavButtonClicked}/>); 
        		})}
            </div>
        );
    }
}