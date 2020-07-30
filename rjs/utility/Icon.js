import React from "react";

export default class Icon extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        return (
            <div className={"no-padding-and-no-margin"} style={{
            display: "inherit",
            backgroundImage: "url('" + this.props.iconPath + "')",
            backgroundPosition: "center",
	        backgroundRepeat: "no-repeat",
	        backgroundSize: "cover",
	        height: this.props.height,
	        width: this.props.width,
        	}}></div>
        );
    }
}