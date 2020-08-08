import React from "react";

import { NavLink } from 'react-router-dom';

export default class GenericNavButton extends React.Component {
    render() {
    	const isActive = (path, match, location) => !!(match || path === location.pathname);
        return (
            <NavLink to={this.props.description["route_path"]} activeClassName="nav-button active"
                className="nav-button"
                class="nav-button"
                style={(this.props.description["text"] != null && this.props.description["text"].length >= 4) ? {width: "100px"} : {}}
                isActive={isActive.bind(this, this.props.description["route_path"])}>

                {this.props.description["text"]}

                <div style={this.props.description["icon"] == null ? {display: "none"} : {
                			height: "100%",
                			width: "25px",
				    		backgroundImage: 'url(' + this.props.description["icon"] + ')',
				            backgroundPosition: "center",
					        backgroundRepeat: "no-repeat",
					        backgroundSize: "cover"}
				}></div>
            </NavLink>
        );
    }
}