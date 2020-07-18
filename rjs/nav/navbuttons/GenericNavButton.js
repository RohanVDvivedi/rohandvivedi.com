import React from "react";

export default class GenericNavButton extends React.Component {
	handleClick() {
		this.props.ifNavButtonClicked(this.props.name)
	}
	handleHover() {
		console.log(this.props.name);
	}
    render() {
        return (
            <a  id={this.props.name + "-nav-button"}
                class={"nav-button" + ((this.props.selected == this.props.name) ? " active" : "")}
                onClick={this.handleClick.bind(this)}>
                    {this.props.name}
            </a>
        );
    }
}