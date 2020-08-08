import React from "react";

import { Link } from 'react-router-dom';

export default class GenericNavButton extends React.Component {
    render() {
        return (
            <Link to={this.props.description["route_path"]} id={this.props.name + "-nav-button"}
                class={"nav-button" + (this.props.isSelected ? " active" : "")}
                style={(this.props.description["text"] != null && this.props.description["text"].length >= 4) ? {width: "100px"} : {}}>

                {this.props.description["text"]}

                <div style={this.props.description["icon"] == null ? {display: "none"} : {
                			height: "100%",
                			width: "25px",
				    		backgroundImage: 'url(' + this.props.description["icon"] + ')',
				            backgroundPosition: "center",
					        backgroundRepeat: "no-repeat",
					        backgroundSize: "cover"}
				}></div>
            </Link>
        );
    }
}