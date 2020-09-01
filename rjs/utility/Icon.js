import React from "react";

export default class Icon extends React.Component {
    render() {
        return (
        	<a class={"generic-content-box-hovering-emboss-border" + ((this.props.infoBoxText != null) ? " tooltip-container" : "") }
        		href={this.props.path} target={this.props.path.includes("#") ? "_self" : "_blank"} 
        		style={{display: "block",padding:this.props.padding}}>
		            <div className={"no-padding-and-no-margin"} style={{
		            backgroundImage: "url('" + this.props.iconPath + "')",
		            backgroundPosition: "center",
			        backgroundRepeat: "no-repeat",
			        backgroundSize: "cover",
			        height: this.props.height,
			        width: this.props.width,
		        	}}></div>
		        	{
		        		(this.props.infoBoxText != null) ? 
		        			(<div class="tooltip-content">{this.props.infoBoxText}</div>) : ""
		        	}
	        </a>
        );
    }
}