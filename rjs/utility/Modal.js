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
    render() {
        return (
        	<div>
        		<div class="modal-background" style = {{
        			display: (this.state.modal_hidden ? "none" : "block"),
        		}}>
        			<div class="modal content-root-background">
        				<div class="modal-header flex-row-container">
	        				<div class="modal-title">
	        					{this.props.title}
	        				</div>
	        				<Icon onClick={this.closeButtonClicked.bind(this)} iconPath="/icon/close.png" infoBoxText="close" height="35px" width="35px" padding="5px" />
	        			</div>
	        			<div class="modal-content">
        					{this.props.children}
        				</div>
        			</div>
        		</div>
        	</div>
        );
    }
}