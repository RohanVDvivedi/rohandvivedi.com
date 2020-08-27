import React from "react";

import { NavLink } from 'react-router-dom';

export default class GenericNavButton extends React.Component {
    render() {
    	var defaultClass = "nav-button";
    	var hasToolTip = false
    	if(this.props.description["pop_up_info"] != null && this.props.description["pop_up_info"].length > 0){
    		defaultClass += " tooltip-container"
    		hasToolTip = true
    	}
        return (
            <NavLink to={this.props.description["route_path"]} activeClassName={defaultClass + " active"}
                className={defaultClass}
                class={defaultClass}
                style={(this.props.description["text"] != null && this.props.description["text"].length >= 4) ? {minWidth: "100px"} : {}}>

                <div>{this.props.description["text"]}</div>

                <div style={this.props.description["icon"] == null ? {display: "none"} : {
                			height: "100%",
                			width: "25px",
				    		backgroundImage: 'url(' + this.props.description["icon"] + ')',
				            backgroundPosition: "center",
					        backgroundRepeat: "no-repeat",
					        backgroundSize: "cover"}
				}></div>

				{hasToolTip ? (<div class="tooltip-content">{this.props.description["pop_up_info"]}</div>) : ""}
            </NavLink>
        );
    }
}