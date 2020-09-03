import React from "react";

import Icon from "./Icon.js"

export default class Modal extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			modal_hidden: true
		}
	}
	closeButtonClicked() {
		this.setState({
			modal_hidden: true
		})
	}
	show() {
		this.setState({
			modal_hidden: false
		})
	}
    render() {
        return (
        	<div style={{display: "contents"}}>
        		<a class="hover-pointer" href="#" onClick={this.show.bind(this)}>
        			{this.props.link}
        		</a>

        		<div class="modal-background" style = {{
        			display: (this.state.modal_hidden ? "none" : "flex"),
        		}}>
        			<div class="modal set_sub_content_background_color generic-content-box-border">
        				<div class="modal-header flex-row-container">
	        				<div class="modal-title">
	        					{this.props.title}
	        				</div>
	        				<Icon onClick={this.closeButtonClicked.bind(this)} iconPath="/icon/close.png" infoBoxText="close" height="30px" width="30px" padding="5px" />
	        			</div>
	        			<div class="modal-content">
        					{this.props.children}
        				</div>
        			</div>
        		</div>
        	</div>);
    }
}