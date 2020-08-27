import React from "react";

export default class Icon extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
    	var showToolTipInfo = false;
    	if(this.props.infoBoxText != null) {
    		showToolTipInfo = true
    	}
        return (
        	<a class={"generic-content-box-hovering-emboss-border" + (showToolTipInfo ? " tooltip-container" : "") }
        		href={this.props.path} target={this.props.path.includes("#") ? "_self" : "_blank"} 
        		style={{display: "block",padding:this.props.padding}}>
		            <div className={"no-padding-and-no-margin"} style={{
		            display: "inherit",
		            backgroundImage: "url('" + this.props.iconPath + "')",
		            backgroundPosition: "center",
			        backgroundRepeat: "no-repeat",
			        backgroundSize: "cover",
			        height: this.props.height,
			        width: this.props.width,
		        	}}></div>
		        	{
		        		showToolTipInfo ? (<div class="tooltip-content">
		        				{this.props.infoBoxText}
		        			</div>) : ""
		        	}
	        </a>
        );
    }
}