import React from "react";

export default class GenericNavButton extends React.Component {
	handleClick() {
		this.props.ifNavButtonClicked(this.props.name)
	}
    render() {
        return (
            <a  id={this.props.name + "-nav-button"}
                class={"nav-button" + (this.props.isSelected ? " active" : "")}
                onClick={this.handleClick.bind(this)}
                style={(this.props.buttonDescription["text"] != null && this.props.buttonDescription["text"].length >= 4) ? {width: "100px"} : {}}>

                {this.props.buttonDescription["text"]}

                <div style={this.props.buttonDescription["icon"] == null ? {display: "none"} : {
                			height: "100%",
                			width: "25px",
				    		backgroundImage: 'url(' + this.props.buttonDescription["icon"] + ')',
				            backgroundPosition: "center",
					        backgroundRepeat: "no-repeat",
					        backgroundSize: "cover"}
				}></div>
            </a>
        );
    }
}